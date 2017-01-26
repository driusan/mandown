# mandown - a tool for writing man pages in markdown.

## Synopsis

mandown [options] filename [filenames..]

## Description

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

mandown will attempt to guess the section of the manual page as best as it can
from context. Files named foo.n.md are assumed to be intended to be
documentation for foo(n). Otherwise, mandown will assume you're documenting a
command and assume section 1.

You need to take a little care when writing your markdown. Your file should
start with a header of the form "# command - title" and each section of the
manpage should start with "## sectionname". Other than that, mandown doesn't
pay much attention to the content of the file and passes most things along
literally.

## Options

-t
	Output the file in troff format, instead of plaintext

--section=n
	specify the section that the manual being parsed belongs to. If unspecified,
	mandown will try and guess based on the filename as described above.

--install=dir
	Copy the gzipped output to `dir` instead of printing it to the screen (implies -t)

## Bugs

This is mostly written as a hack and proof of concept. It probably has many.

Please file any specific issues that you encounter at
https://github.com/driusan/mandown

## Author
Dave MacFarlane <driusan@gmail.com>

## See Also

man(1), groff(1), nroff(1), troff(1)
