# Forge example templates

A starting point for the templates repo of your own Forge. `forge create`
scaffolds new services from a templates repo in your GitHub organization
(`<org>/forge-templates` unless your install names another). This repo shows
the format with one template you can use as it is.

| Template | What you get |
| --- | --- |
| `hello` | A small Go web service. CI runs on GitHub-hosted runners, pushes the image to `ghcr.io/<org>/<service>` with the workflow's own token, and deploys `main` to staging with `forge deploy` |

## Use it

1. Copy this repo's files into `<org>/forge-templates` (an empty repo the
   Forge GitHub App is installed on).
2. In `hello/template.yaml`, set the default of `forgeApi` to your Forge's
   address, for example `https://forge.example.com`. Every new service's CI
   deploys through it. (An organization variable `FORGE_API_URL` overrides
   it, but on GitHub Free private repositories cannot read organization
   variables.)
3. Check the templates, then release them with a tag on the whole repo:

   ```sh
   forge template lint
   git tag v1.0.0 && git push origin v1.0.0
   ```

4. Make a service:

   ```sh
   forge template list
   forge create hello-world --template hello --team <your team>
   ```

Forge pulls the images with the pull secret of your install (`apps.pullSecret`
in the chart's values). The images are private, like the repos `forge create`
makes, so that secret needs a token that can read your organization's
packages on ghcr.io.

## The format

A template is a directory with a `template.yaml` (what it is, the parameters
it asks for) and a `skeleton/` whose files become the new repo. Paths and
contents are Go templates; `hello` uses `{% %}` as delimiters because GitHub
Actions uses `${{ }}`. The values are `.Service` (`Name`, `Team`, `Owners`),
`.Repo` (`Owner`, `Name`, `URL`), `.Template` (`Name`, `Version`) and
`.Params`. Forge's documentation has the full reference.

`forge template lint` checks every template here without a Forge: it renders
each one with sample values and checks the `forge.yaml` it makes. This repo's
CI runs it on every push.
