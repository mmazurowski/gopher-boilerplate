# Project Deploy

**Description:**
Provide current commit tag and deploys service under specified AWS Profile.

**How to execute:**
```shell
$ task project:deploy
```

**Environmental variables:**
```
GIT_TAG:
  sh: git describe --always --tags
```
