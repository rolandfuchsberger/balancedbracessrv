# Balanced Braces Srv

## Requirements

- Validates if an expression has balanced braces - i.e. every opened braces is closed in the inverse order (e.g. `{test()}` is balanced whereas `{test(})` is unbalanced)
- Exposes logic over a REST endpoint
- REST endpoint should include CORS headers
- HTML output
- High unit test coverage
- CI/CD pipeline for automatic lint, test (deployment not included)
- Docker build image

## Special features

- CI/CD pipeline with travis
- Usage of gobuffalo/plush and gobuffalo/packr/v2

## Further Optimizations

- Speed up builds by using cached docker images (see <http://rundef.com/fast-travis-ci-docker-build>) and possibly reduce parallelization.
- Build docker image in CI/CD and push to repo
- Include XSRF headers
- Define and enforce limits on API

## Key Concepts Included

- Usage of Chi router
- Unit tests
- Benchmarks (html/template vs gobuffalo/plush)
- Test suite

## Key Concepts not Included

- Concurrency / go routines / channels / mutex...
- Dependency injection (wire)
- Service listening on signals
