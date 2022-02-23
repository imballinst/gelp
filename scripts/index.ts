import { Octokit } from '@octokit/rest';
import semverIncrement from 'semver/functions/inc';
import semverPrereleaseParse from 'semver/functions/prerelease';

import fs from 'fs/promises';
import path from 'path';

const OWNER = process.env.OWNER || 'imballinst';
const REPO = process.env.REPO || 'gelp';
const FILENAME = process.env.FILENAME || 'gelp-linux-amd64';

async function main() {
  const octokit = new Octokit({ auth: process.env.RELEASE_TOKEN });

  let latestTag = '';
  let releaseId = -1;
  let isDraft = false;
  let previousBody = '';
  let preId: string | undefined = undefined;

  try {
    const response = await octokit.repos.listReleases({
      owner: OWNER,
      repo: REPO,
      page: 1,
      per_page: 1
    });
    releaseId = response.data[0].id;
    latestTag = response.data[0].tag_name;
    isDraft = response.data[0].draft;
    previousBody = response.data[0].body || '';

    const latestPrereleaseId = semverPrereleaseParse(latestTag);
    if (latestPrereleaseId !== null) {
      preId = `${latestPrereleaseId[0]}`;
    } else {
      preId = 'alpha';
    }
  } catch (err) {
    console.error(err);
    latestTag = 'v0.0.0';
  }

  if (isDraft) {
    // If still draft, then update it.
    await octokit.repos.updateRelease({
      release_id: releaseId,
      owner: OWNER,
      repo: REPO,
      tag_name: latestTag,
      prerelease: true,
      draft: true,
      name: latestTag,
      target_commitish: 'main',
      body: previousBody
    });
    const releaseAssets = await octokit.repos.listReleaseAssets({
      owner: OWNER,
      repo: REPO,
      release_id: releaseId
    });

    const gelp = releaseAssets.data.find((entry) => entry.name === FILENAME);

    if (gelp !== undefined) {
      // Delete the asset first before re-uploading it.
      await octokit.repos.deleteReleaseAsset({
        owner: OWNER,
        repo: REPO,
        asset_id: gelp.id
      });
    }

    await octokit.repos.uploadReleaseAsset({
      data: (
        await fs.readFile(path.join(__dirname, `../publish/${FILENAME}`))
      ).toString(),
      name: FILENAME,
      owner: OWNER,
      repo: REPO,
      release_id: releaseId
    });
  } else {
    // If not draft, then we create a new draft prerelease.
    const possiblyNextTag = semverIncrement(
      latestTag,
      'prerelease',
      {
        includePrerelease: true
      },
      preId
    );

    if (possiblyNextTag !== null) {
      latestTag = possiblyNextTag;
    }

    const response = await octokit.repos.createRelease({
      owner: OWNER,
      repo: REPO,
      tag_name: latestTag,
      prerelease: true,
      draft: true,
      name: latestTag,
      target_commitish: 'main',
      body: 'this is a test release'
    });
    await octokit.repos.uploadReleaseAsset({
      data: (
        await fs.readFile(path.join(__dirname, `../publish/${FILENAME}`))
      ).toString(),
      name: FILENAME,
      owner: OWNER,
      repo: REPO,
      release_id: response.data.id
    });
  }
}

main();
