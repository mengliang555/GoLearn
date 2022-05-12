# Golang Test

- Unit Test
- [Fuzz Test](https://go.dev/doc/fuzz/#glossary)
    - tips:
        - for FUZZ function type must have no type parameters
        - The First parameters of Fuzz function must be *testing.T
        - [go test -fuzz=Fuzz] or [go test -fuzz={FuzzTestName}]
        - [-fuzztime 10s(eg)]fuzz execute time, default indefinitely
        - [-parallel] processes running at once, default $GOMAXPROCS. Currently, setting -cpu during fuzzing has no
          effect.
        - 