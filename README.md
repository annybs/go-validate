# Go Validate

A suite of straightforward validation functions. You put something in, you get back `nil` or an error.

## Error handling

You can use `errors.Is()` to ascertain the type of errors returned by validation functions. This may be helpful to control side effects, particularly if you are using multiple validators or want to change the error message.

## License

See [LICENSE.md](../LICENSE.md)
