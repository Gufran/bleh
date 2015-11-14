# Bleh!
Generate some basic boilerplate code for a new Go app

# Installation
```
go get -u github.com/Gufran/bleh
```
bleh!

## Usage
```
bleh
```
bleh!

## WTF?
`bleh` generates a bare minimum skeleton for a Go app. It will create a `Makefile`, `.gitignore`, `.travis.yml`
along with some build specific information about your app, and configure the build step in makefile to populate
that build specific information.
It will also initialise the git repository for you, add a remote and perform an initial commit.
You can instruct `bleh` to not generate that initial commit by including `-no-commit` flag.

That is pretty much all to it, I'm sorry if you were expecting unicorns.

## Contributing
Although I doubt you can add anything else to this tool anymore, I'd love to see what you have to offer.
Fork the repository, make your changes and send a pull request.

## License

The MIT License (MIT)

Copyright (c) 2015 Mohammad Gufran

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

