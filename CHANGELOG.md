# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## v0.3.2 - 2021-08-21
### Added
* Added go 1.17 support.

## v0.3.1 - 2021-02-21
### Added
* Added Go 1.16 support.

## v0.3.0 - 2021-02-06
### Changed
* **BREAKING**: Renamed ```RandomCryptoBytes``` function to ```CryptoBytes```.
* **BREAKING**: Renamed ```RandomCryptoString``` function to ```CryptoString```.
* **BREAKING**: Renamed ```RandomMathBytes``` function to ```MathBytes```.
* **BREAKING**: Renamed ```RandomMathString``` function to ```MathBytes```.
* **BREAKING**: Renamed ```GenerateUUID``` function to ```UUIDV4```.

## v0.2.2 - 2020-12-31
### Changed
* Updated readme.

## v0.2.1 - 2020-12-30
### Added
* Added GitHub Stale action.

## v0.2.0 - 2020-11-28
## Added
* Added Go 1.15 support.
* Added ```RandomCryptoString``` and ```RandomMathString``` functions.

### Changed
* **BREAKING**: Renamed package name from ```rand``` to ```xrand```.
* **BREAKING**: Renamed ```GenerateCryptoBytes``` function to ```RandomCryptoBytes```.
* **BREAKING**: Renamed ```GenerateCryptoString``` function to ```RandomCryptoString```.

## v0.1.0 - 2020-06-11
### Added
* Added ```GenerateCryptoBytes```, ```GenerateCryptoString``` ```GenerateUUID``` functions.