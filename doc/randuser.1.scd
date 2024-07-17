randuser(1)

# NAME

randuser - generate random usernames

# SYNOPSIS

*randuser* [options...]

# DESCRIPTION

*randuser* is a command-line application that generates random usernames based
on a specified format. It combines adjectives, verbs, nouns, and numbers to
create unique and memorable usernames using the same algorithm as diceware
passwords. By default, the "{{ adjective }}{{ noun }}" format is used.

# OPTIONS

*-t*, *--title-case*
	Generate the username using title case (each word capitalized).

*-f* _FORMAT_, *--format* _FORMAT_
	Specify the format of the generated username. The default format is "{{
	adjective }}{{ noun }}".

	The following placeholders are available:

	*{{ adjective }}*
		Insert a random adjective.

	*{{ verb }}*
		Insert a random verb.

	*{{ noun }}*
		Insert a random noun.

	*{{ number }}*
		Insert a random number.

*-h*, *--help*
	Show help message and quit.

*-v*, *--version*
	Show version information and quit.

# FORMAT

The format string specifies how the username should be constructed. It can
contain any combination of the available placeholders. Spaces and static text in
the format string are ignored.

Examples of valid formats:

- "{{ adjective }}{{ noun }}"
- "{{ verb }}{{ noun }}{{ number }}"
- "{{ adjective }}\_{{ noun }}"

# EXAMPLES

*Example 1. Generate a random username with default format:*
	$ randuser

*Example 2. Generate a random username with title case:*
	$ randuser -t

*Example 3. Generate a random username with custom format:*
	$ randuser -f '{{ verb }}{{ noun }}{{ number }}'

*Example 4. Generate a random username with custom format and title case:*
	$ randuser -t -f '{{ verb }}{{ noun }}{{ number }}'

# AUTHORS

Maintained by James Pond <james@cipher.host>.
