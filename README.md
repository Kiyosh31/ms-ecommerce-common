# Overview

`Common` library for go microservices arch project [ms-ecommerce](https://github.com/Kiyosh31/ms-ecommerce)

# How to push a new commit?

I have automated the tag versioning and push, the only thing to do to push a new commit is:

```console
make push
```

This will fetch the last tag in GH and then increment it +1 and push it
