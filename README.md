# Am I Pwned?

Uses [';--have i been pwned?](https://haveibeenpwned.com) dataset to verify if your passwords have been compromised.

# Installation

    go get -u github.com/jcgay/am-i-pwned
    
# Usage

## Verify a password

    $> am-i-pwned password [candidate]
    
or type your password interactively using:

    $> am-i-pwned password
    
A SHA1 fragment of your password will be send to `https://haveibeenpwned.com` API and will be compared locally with service response. See [Searching a password by range](https://haveibeenpwned.com/API/v2#SearchingPwnedPasswordsByRange) for more information.

The number of occurence of the password in the database will be print in the console, if this number is 0, no entry has been found.

## Check your LastPass account

    $> am-i-pwned lastpass
    
All your passwords will be checked (same method as [single password check](#verify-a-password)) using [LastPass CLI](https://github.com/lastpass/lastpass-cli).  
You'll have to install [LastPass CLI](https://github.com/lastpass/lastpass-cli) and be logged to use this command.

## Help

    $> am-i-pwned help [command]