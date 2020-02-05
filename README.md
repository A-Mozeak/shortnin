# Shortnin'
A link shortening API written in Go.

# Getting Started

## Installation
1. Install Go
1. Run ```go build```
1. Run ```./shorty```

## Usage
### *Create*
To generate a link, send a **POST** request to ```/create?{params}``` with the following parameters:
* url - The URL to be shortened.
* custom (optional) - An optional string for a custom short URL.

### *Stats*
To view link usage statistics, send a **GET** request to ```/stats?{params}``` with the following parameters:
* link - The short link to be viewed.

# Design
## Language/Frameworks
I have done most of my professional work using Node.js, but I felt that Go would be at least as good a fit for this project, if not better. The reasons why I chose Go are:
- Tooling
  - Go offers a great set of tools out of the box. Dependency management is a breeze, thanks to ```go get```, and being able to pull in parts of the standard library automatically as I need them helps me to be more productive.
- Standard Library
  - The Go standard library has a deep set of tools for working with http, json, and other web technologies.
- Syntax
  - Go has a nice, clear, terse syntax that makes it easier for me to read through different chunks as I need. This is also reflected in the handling of modules, since capitalization denotes a package-level export.
- Challenge
  - I only picked up Go in December, and I want to demonstrate that I have a grasp of web concepts that transcends any particular language.

I also used **Mux** to handle the app routing. Mux is a popular tool among Go developers because it abstracts away some of the low-level specifics of the ```http.ServeMux``` in the standard library. I also appreciate that it makes the routing component of this project a great deal easier to read.

## Storage
The MockDB struct is a mock database that consists of only two maps:
1. A map from long urls to randomly-generated shortlinks. I'm using this map to check if a shortlink has already been generated for the given url.
1. A map from shortlinks to Shorty structs. I'm using this map to address the data associated with a given shortlink. This structure also allows us to map multiple custom shortlinks to the same long URL.

The Shorty struct is a data model that associates a shortlink with a long URL, and keeps track of the shortlinks usage. Each Shorty is serializable into a JSON object.

## Generating the Shortlinks
For this exercise, I went with something pretty brute-force for generating the shortlinks.
1. Seed a random number generator with the current Unix time.
1. Establish an alphabet (in this case [a-zA-Z0-9]).
1. For 6 turns, use the RNG to choose one value, by index, from the alphabet and concatenate it to a string (I used stringBuilder to minimize copying the string).
1. Return the generated string.