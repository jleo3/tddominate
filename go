---
layout: post
title: "G"
date: 2013-04-02 16:29
comments: true
categories: 
---

Why are there Go Programmers? - Blake Mizerany

Want to know what Go programmers look like? Here's one:
[shahid pic]

Look! Here's another one:
[dix pic]

Coming into this talk all I know about Go is that a couple of people I trust have found it suitable for their recent production applications. Here I'm hoping to learn much more. After all, there's still time for me to abandon this whole "learning Scala" thing (wink).

So far I like the idea of the language being built by an impressive team of engineers and it's promise of super-fast compile times. Blake hasn't been this excited about a language in over a decade. So that's something.

Go is defined in UTF-8 and thus supports it 100%. Cool things:

* Benchmarking tools
* Concurrency primitives
* Functional aspects
* Fast as a feature
* Simplicity
* Testing tools
* Statically linked binaries mean that everything compiles down to machine code, not to a vm. 

This presentation software is a Go program! Well, holy crap. And I just found out that I am need of some presentation software!
godoc.org/code.google.com/p/go.talks/pkg/present

go fmt - the answer to code formatting arguments. Just run it against your source and then check the code in. go fmt will format your code according to the standard.

The testing library seams very cool in that it regression tests the examples in one's documentation. This can keep you sane and keep your APIs honest. 

The language is easy to follow, terse but expressive. The approach (at least in Blake's example) reminds me of the virtues extolled in Clojure. The minimal syntax also gives Go a functional feel. 

I like Blake's stance on Go, coming from the context as the creator of Sinatra. He says that he used to love DSLs and beautiful code but now has a greater love for simplicity and speed. It's interesting (and accurate) that these principles can be trade-offs for one another.

Where is Go going? One goal is for its HTTP library to be the most complete HTTP library. Goroutines are one way that HTTP programming is effective in Go. Goroutines are lightweight threads that start small and grow and shrink as needed. These are managed by the Go language, which means millions of them could be allocated under the hood depending on memory and compute resources of your system. But the Go programmer doesn't have to worry about it - just run your program and the goroutines will manage themselves. 

This looks like an interesting language surrounded by an intellectual community. I'm going to do some more poking around and see if I get myself lured in.
