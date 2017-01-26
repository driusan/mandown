# mandown - a tool for writing man pages in markdown.

mandown converts markdown formatted manuals into the troff(1) based format that
is used by man(1) on most systems.

Man pages are useful, but often neglected by modern software developers. This
may be because man pages need to be written in a typesetting language seldom
used by most developers (outside of man), or it may be that the software
being written is written in a language that can be used on non-UNIX systems
and the effort of writing extra documentation for a subset of users isn't
perceived as worth the effort. 

Writing the documentation in markdown addresses both of these problems. Modern
software developers are already likely familiar with it, and it's consumable
without any special software on non-UNIX systems.

By default, mandown will take any markdown file passed to it, try to parse it,
normalize the sections/whitespace, and print it to the screen in a format
similar to running `man <cmd>`. With the -t option, it will print it as troff
macros instead of human-readable format.

See mandown.1.md for an example.

## Installation

mandown is written in Go, and should be go-gettable

```
go get github.com/driusan/mandown
```

