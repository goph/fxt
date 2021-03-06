# Change Log


All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).


## [Unreleased]


## [0.30.0] - 2018-05-12

### Added

- Default app modules for http and grpc
- Default app runners for http and grpc

### Changed

- Moved gorilla tracing to a separate package
- Renamed `context` package to `fxcontext`
- Renamed `daemon` package to `fxdaemon`
- Renamed `database/sql` package to `fxsql`
- Renamed `debug` package to `fxdebug`
- Renamed `debug/metrics/prometheus` package to `fxpromhttp`
- Renamed `dev` package to `fxdev`
- Renamed `errors` package to `fxerrors`
- Renamed `errors/airbrake` package to `fxairbrake`
- Renamed `grpc` package to `fxgrpc`
- Renamed `http` package to `fxhttp`
- Renamed `http/gorilla` package to `fxmux` and moved to `http/gorilla/mux`
- Renamed `http/gorilla/opentracing` package to `otmux` and moved to `http/gorilla/mux/opentracing`
- Renamed `log` package to `fxlog`
- Renamed `log/logrus` package to `fxlogrus`
- Renamed `prometheus` package to `fxprometheus`
- Renamed `test` package to `fxtesting`
- Moved `tracing` package to `tracing/opentracing` under the name `fxopentracing`
- Renamed `tracing/jaeger` package to `fxjaeger`
- Renamed `tracing/jaeger/metrics/prometheus` package to `jaegerprometheus`
- Moved `debug/metrics/prometheus` package to `metrics/prometheus/promhttp` as `fxpromhttp`
- Moved param and result structs next to the constructor
- Use new Jaeger config constructor instead of the deprecated one

### Removed

- Unused jaeger client config


## [0.23.0] - 2018-04-29

### Added

- http: opentracing correlation ID middleware

### Changed

- http: rewritten correlation ID middleware
- grpc: rewritten correlation ID middleware


## [0.22.0] - 2018-04-26

### Added

- context: New package for basic context keys
- http: correlation ID middleware

### Changed

- Move random string generation to internal package for reusability
- Move common Correlation ID related interfaces to `internal/correlationid` package
- Use `context.Context` wherever possible


## [0.21.0] - 2018-04-26

### Added

- log: [logrus](https://github.com/sirupsen/logrus) adapter


## [0.20.0] - 2018-03-18

### Added

- metrics/prometheus: Prometheus instance
- tracing/jaeger: Prometheus metrics factory
- Component log context to prometheus and jaeger

### Changed

- metrics/prometheus: Moved prometheus handler to debug package


## [0.19.1] - 2018-03-05

### Added

- database/sql: Comments for DB connection config

### Fixed

- database/sql: Invalid dependency injection container parameter


## [0.19.0] - 2018-02-28

### Fixed

- database/sql: max idle connection setting falls back to the internal default by default


## [0.18.0] - 2018-02-28

### Added

- database/sql: DB connection configuration
- database/sql: Utility function to make the master connection the primary one

### Fixed

- database/sql: Slave overwriting master connection


## [0.17.0] - 2018-02-04

### Added

- Test runner interface for test runners (already implemented by Godog and Main test runners)
- Test runner list for merging test runners into one
- Test runner factories
- Function type for test runner factories

### Changed

- gRPC client acceptance context moved to subpackage to avoid dependency

### Removed

- AppContext function


## [0.16.0] - 2018-02-01

### Added

- gRPC Dial function
- DB acceptance context
- gRPC client acceptance context
- App acceptance context


## [0.15.1] - 2018-01-24

### Added

- Error indicators for http and debug server shutdown


## [0.15.0] - 2018-01-24

### Added

- Invoke function for configuring global airbrake logger
- Application info struct for common application details

### Fixed

- Airbrake ProjectKey configuration option


## [0.14.0] - 2018-01-23

### Added

- Application context function for Godog
- Environment variable loading helper


## [0.13.0] - 2018-01-23

### Changed

- Move `fxtest` package under `test`


## [0.12.0] - 2018-01-23

### Added

- `fx.App` wrapper so that closer does not have to be manually populated
- `fxtest.App` wrapper so that closer does not have to be manually populated

### Fixed

- Segmentation fault when Godog runner used without initialization


## [0.11.0] - 2018-01-22

### Added

- [Godog](https://github.com/DATA-DOG/godog) runner configuration
- Build tag helpers to determine whether a test is executed with `unit`, `acceptance` or `integration` tags
- `GetFreePort` for getting a free TCP port (useful for acceptance tests involving a server connection)


## [0.10.0] - 2018-01-18

### Removed

- TLS option for DB app config


## [0.9.0] - 2018-01-16

### Added

- gRPC correlation ID interceptor


## [0.8.0] - 2018-01-14

### Added

- Metrics support for Jaeger

### Changed

- Upgrade Jaeger libraries
- Use go-kit logger wrapper from jaeger lib
- Upgraded gRPC


## [0.7.0] - 2018-01-12

### Added

- App config for SQL
- App config for Airbrake
- App config for Jaeger
- App config for gRPC client


## [0.6.3] - 2017-12-05

### Added

- `io.Writer` adapter for go-kit logger

### Changed

- Use `io.Writer` log adapter instead of the stdlib one in grpc log


## [0.6.2] - 2017-11-17

### Added

- Error handler result for error handler groups


## [0.6.1] - 2017-11-16

### Added

- StatusChecker health check constructor


## [0.6.0] - 2017-11-16

## Changed

- Upgrade [go.uber.org/dig](http://go.uber.org/dig) to minimum 1.2.0
- Use dig 1.2 features in error handling


## [0.5.2] - 2017-10-30

### Changed

- Updated dependencies


## [0.5.1] - 2017-10-26

### Added

- Stack trace to some errors

### Changed

- Updated dependencies


## [0.5.0] - 2017-10-18

### Added

- Prometheus metrics endpoint register function
- Custom gRPC server options to the gRPC config

### Changed

- Moved Gorilla Mux router to `http/gorilla` package
- Replaced custom router solution with tracer injection

### Removed

- Custom Mux router wrapper
- Interceptors from gRPC constructor params


## [0.4.0] - 2017-10-15

### Added

- Router wrapping Gorilla mux adding OpenTracing support

### Changed

- Logger is now optional in Jaeger tracer


## [0.3.3] - 2017-10-13

### Fixed

- Null pointer closer


## [0.3.2] - 2017-10-13

### Fixed

- Null pointer closer


## [0.3.1] - 2017-10-13

### Changed

- Lock go-kit to patch versions
- Register the daemon runner in the application lifecycle


## [0.3.0] - 2017-10-13

### Added

- Daemon and cron tools
- `WithTimeout` function to enforce timeouts
- Simple http server construction
- Database connection constructor
- Master-slave database connection constructor
- Add gRPC server constructor

### Changed

- Do not require a `net.Addr` in debug server config
- Debug server is started without goph/serverz
- Make configurations non-optional options


## [0.2.0] - 2017-10-12

### Added

- Error channel returned from debug server
- Opentracing global tracer

### Changed

- Provide debug.Handler from debug server constructor
- Make logger optional (but highly recommended) for debug server

### Removed

- Global error channel
- Debug bootstrap
- goph/stdlib requirement


## 0.1.0 - 2017-10-12

### Added

- [go-kit](https://github.com/go-kit/kit/tree/master/log) logger constructor
- [emperror](https://github.com/goph/emperror) error handler constructor
- [airbrake](https://github.com/airbrake/gobrake) compatible error handler constructor
- Error handler stack for collecting multiple handlers under a single dependency
- [opentracing](http://opentracing.io/) tracer constructor
- [jaeger](https://github.com/jaegertracing) tracer constructor
- Lifecycle extension for supporting Closers
- Debug server constructor (invoke fn)
- Health collector constructor
- Bootstrap options


[Unreleased]: https://github.com/goph/fxt/compare/v0.30.0...HEAD
[0.30.0]: https://github.com/goph/fxt/compare/v0.23.0...v0.30.0
[0.23.0]: https://github.com/goph/fxt/compare/v0.22.0...v0.23.0
[0.22.0]: https://github.com/goph/fxt/compare/v0.21.0...v0.22.0
[0.21.0]: https://github.com/goph/fxt/compare/v0.20.0...v0.21.0
[0.20.0]: https://github.com/goph/fxt/compare/v0.19.1...v0.20.0
[0.19.1]: https://github.com/goph/fxt/compare/v0.19.0...v0.19.1
[0.19.0]: https://github.com/goph/fxt/compare/v0.18.0...v0.19.0
[0.18.0]: https://github.com/goph/fxt/compare/v0.17.0...v0.18.0
[0.17.0]: https://github.com/goph/fxt/compare/v0.16.0...v0.17.0
[0.16.0]: https://github.com/goph/fxt/compare/v0.15.1...v0.16.0
[0.15.1]: https://github.com/goph/fxt/compare/v0.15.0...v0.15.1
[0.15.0]: https://github.com/goph/fxt/compare/v0.14.0...v0.15.0
[0.14.0]: https://github.com/goph/fxt/compare/v0.13.0...v0.14.0
[0.13.0]: https://github.com/goph/fxt/compare/v0.12.0...v0.13.0
[0.12.0]: https://github.com/goph/fxt/compare/v0.11.0...v0.12.0
[0.11.0]: https://github.com/goph/fxt/compare/v0.10.0...v0.11.0
[0.10.0]: https://github.com/goph/fxt/compare/v0.9.0...v0.10.0
[0.9.0]: https://github.com/goph/fxt/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/goph/fxt/compare/v0.7.0...v0.8.0
[0.7.0]: https://github.com/goph/fxt/compare/v0.6.3...v0.7.0
[0.6.3]: https://github.com/goph/fxt/compare/v0.6.2...v0.6.3
[0.6.2]: https://github.com/goph/fxt/compare/v0.6.1...v0.6.2
[0.6.1]: https://github.com/goph/fxt/compare/v0.6.0...v0.6.1
[0.6.0]: https://github.com/goph/fxt/compare/v0.5.2...v0.6.0
[0.5.2]: https://github.com/goph/fxt/compare/v0.5.1...v0.5.2
[0.5.1]: https://github.com/goph/fxt/compare/v0.5.0...v0.5.1
[0.5.0]: https://github.com/goph/fxt/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/goph/fxt/compare/v0.3.3...v0.4.0
[0.3.3]: https://github.com/goph/fxt/compare/v0.3.2...v0.3.3
[0.3.2]: https://github.com/goph/fxt/compare/v0.3.1...v0.3.2
[0.3.1]: https://github.com/goph/fxt/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/goph/fxt/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/goph/fxt/compare/v0.1.0...v0.2.0
