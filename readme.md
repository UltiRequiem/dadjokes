# Dadjokes

[![Go Reference](https://pkg.go.dev/badge/github.com/UltiRequiem/dadjokes/pkg.svg)](https://pkg.go.dev/github.com/UltiRequiem/dadjokes/pkg)

Get random dad jokes using the
[icanhazdadjoke API](https://icanhazdadjoke.com/api) 👨

### Install

```sh
go get github.com/UltiRequiem/dadjokes/pkg
```

### Usage

```go
import (
	"fmt"
	"github.com/UltiRequiem/dadjokes/pkg"
)

func main() {
	joke, errorGettingJoke := dadjokes.GetJoke()

	if errorGettingJoke != nil {
		panic(errorGettingJoke)
	}

	fmt.Println(joke.Joke)
}
```

### Documentation

Is hosted on
[pkg.go.dev](https://pkg.go.dev/github.com/UltiRequiem/dadjokes/pkg) 📖

## CLI

### Installation

```sh
go install github.com/UltiRequiem/dadjokes@latest
```

### Usage

```sh
dadjokes
```

Or 👇

```sh
dadjokes 15
```

![Screenshot](https://user-images.githubusercontent.com/71897736/162102151-2c141e36-9f66-48ac-8b8a-d5b371498f73.png)

[Video Showcase](https://youtu.be/eo-1zpdiLnY) 📹

## Support

Open an Issue, I will check it a soon as possible 👀

If you want to hurry me up a bit
[send me a tweet](https://twitter.com/UltiRequiem) 😆

Consider [supporting me on Patreon](https://patreon.com/UltiRequiem) if you like
my work 🙏

Don't forget to start the repo ⭐

## Versioning

We use [Semantic Versioning](http://semver.org). For the versions available, see
the [tags](https://github.com/UltiRequiem/dadjokes/tags) 🏷️

## Authors

[Eliaz Bobadilla](https://ultirequiem.com) - Creator and Maintainer 💪

See also the full list of
[contributors](https://github.com/UltiRequiem/dadjokes/contributors) who
participated in this project ✨

## Licence

Licensed under the MIT License 📄
