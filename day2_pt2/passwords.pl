#!/bin/env perl
 
use strict;
use warnings;
use feature qw(say);

# While it appears you validated the passwords correctly, they don't seem to be what the Official Toboggan Corporate Authentication System is expecting.

# The shopkeeper suddenly realizes that he just accidentally explained the password policy rules from his old job at the sled rental place down the street!
# The Official Toboggan Corporate Policy actually works a little differently.

# Each policy actually describes two positions in the password, where 1 means the first character, 2 means the second character, and so on.
# (Be careful; Toboggan Corporate Policies have no concept of "index zero"!) Exactly one of these positions must contain the given letter. 
# Other occurrences of the letter are irrelevant for the purposes of policy enforcement.

# Given the same example list from above:

# 1-3 a: abcde is valid: position 1 contains a and position 3 does not.
# 1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
# 2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.

# How many passwords are valid according to the new interpretation of the policies?

my $password_file = "passwords";

chomp(my @lines = <STDIN>);

my $count = 0;

foreach my $line (@lines) {
	if($line =~ m/^(\d+)-(\d+) ([A-Za-z])\: ([A-Za-z]+)$/) {
		my $first_position = $1;
		my $second_position = $2;
        my $target = $3;
        my $password = $4;

        # print $first_position;
        # print $second_position;
        # print $target;
        # print $password;
        # print "\n";

        my @pwds = split('', $password);

        # print "Str: $line Targ: $target First: $first_position Second: $second_position Pass: $password 1st: $pwds[($first_position - 1)] 2nd: $pwds[($second_position - 1)]\n";
        # print "Str: $line 1st: $pwds[($first_position - 1)] 2nd: $pwds[($second_position - 1)]\n";
        # if(($pwds[$first_position - 1] eq $target) || ($pwds[$second_position - 1] eq $target)) {
        if($pwds[$first_position - 1] eq $target) {
            print "Str: $line 1st: $pwds[($first_position - 1)]\n";
            $count++;
            last;
        }
        elsif ($pwds[$second_position - 1] eq $target)  {
            print "Str: $line 2nd: $pwds[($second_position - 1)]\n";
            $count++;
            last;
        }
	}
}

print $count, "\n";


# open(my $fh, '<', $password_file) or die "Could not open the damn password file $password_file $!";

# while (my $row = <$fh>) {
#     chomp $row;
#     print "$row\n";
# }
