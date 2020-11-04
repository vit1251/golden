# Manual

This manual detail review most popular Golden Point use case.

## Initial setup

1. Start Golden Point executable
2. Open your Web-browser address http://127.0.0.1:8080/setup

### Detailed description of parameters in the settings section

Parameter descriptions contain a short definition, as well as an example enclosed in quotation marks (quotation marks indicate the boundaries of the value and are not part of it).

#### RealName

The *RealName* field contains the real username used in the correspondence.

Example: "Ivan Petrov"

#### Origin

This line appears near the bottom of a message and gives a small amount of information about the system where it originated.

Note: In the Origin field, you can specify the path to a file with several lines written in UTF-8 encoding.
      In this case, a random line will be choice from this file.
      The path must prefixed by the "@" prefix

Example: "@C:\Users\vit12\Fido\Origin.txt"
Example: "The Conference Mail BBS"
