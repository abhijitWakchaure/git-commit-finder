# git-commit-finder

This tool will help you find the latest git commit ID / SHA 256 hash for given repo where the files in the directory matches exactly to the git repo.

## Where to use this tool

Use this tool when you have a old go dependency vendored somewhere but don't know the exact version of the dependency. Just clone the git repo for your dependency and run this tool, this will get you the exact commit ID for your vendors.

## How to use

Just install this tool using command:

```bash
go install github.com/abhijitWakchaure/git-commit-finder@latest
```

Then you can use this tool like this:

```bash
$ git-commit-finder --dir /home/abhijit/test/dateparse-copy --gitRepo /home/abhijit/test/dateparse
Starting git-commit-finder v1.0.0
Scanning directory: /home/abhijit/test/dateparse-copy
Found 250 commits for your repo: /home/abhijit/test/dateparse
1/250 Checked out commit ID 6b43995a97dee4b2c7fc0bdff8e124da9f31a57e
2/250 Checked out commit ID 5dd51ed0f76a3790e35b502070d9acbb404fab30
3/250 Checked out commit ID 0eec95c9db7e40737864608615f5aff4cac78680
...
...
...
16/250 Checked out commit ID fcfe3a02eb309c5ec8c1d9b5da30c217663c30e2

Found the matching commit: fcfe3a02eb309c5ec8c1d9b5da30c217663c30e2
```
