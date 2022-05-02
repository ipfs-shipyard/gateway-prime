# Gateway Prime

**This an extraction of the [`corehttp`](https://github.com/ipfs/go-ipfs/tree/master/core/corehttp) module from `go-ipfs`,
abstracted to only require a narrow API interface to back the HTTP serving code.**

The backing API provided is defined in `api.go`, and uses an ipld linksystem, and a fetcher for loading additional requested data.

## License & Copyright

Copyright &copy; 2022 Protocol Labs

Licensed under either of

 * Apache 2.0, ([LICENSE-APACHE](LICENSE-APACHE) / http://www.apache.org/licenses/LICENSE-2.0)
 * MIT ([LICENSE-MIT](LICENSE-MIT) / http://opensource.org/licenses/MIT)

### Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in the work by you, as defined in the Apache-2.0 license, shall be dual licensed as above, without any additional terms or conditions.
