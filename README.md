# MassPostage
A simple command line mail merge app written in Go.

`MassPostage` uses plain text files and CSV files to generate personalized email files that can later be sent using regular email apps.

**Table of Contents**
- [Quickstart](#quickstart)
- [Usage](#usage)
- [Project status](#project-status)
- [Acknowledgements](#acknowledgements)
- [Authors](#authors)
- [License](#license)

## Quickstart

You need to clone the repository and compile the code to get started. Run the following commands to get a working version:

```bash
git clone https://github.com/IvanIsCoding/MassPostage
cd MassPostage
make clean
make build
```

A binary will be generated in the `bin` folder. Usage bellow assumes that the binary is called `masspost`.

## Usage

MassPostage takes three main arguments: the CSV file, the email body and the campaign name.

>     masspost <csv file> <email body> <campaign name>
>     masspost example.csv example.eml ExampleCampaign

The CSV file contains the data that will be replaced in the email body to generate personalized emails.

The email body should be a text file with the content of the message. To use the content from each line of the CSV, insert the exact name of the column inside double curly brackets like `Column Name` (disclaimer: column names are case sensitive). An example of a file would be:

```
TO: {{Email Adddress}}
SUBJECT: Example Campaign
FROM: Foo Bar <foo@bar.com>

Hi {{Name}},

This is an example email.

Cheers,

Foo Bar
```

At the end of the execution, personalized EML files will be generated in the campaign folder for each line of the CSV. Afterwards, you can open those files with your favorite email client and send the emails.

## Project status

The project is still a work in progress. Here is the current status of features of the app.

- [x] Mass mail generator for simple text files
- [ ] Support attachments in mail generator  
- [ ] Mass mail generator for Markdown files
- [ ] Mail sender for generated campaigns
- [ ] Write tests for the codebase

## Authors

* **Ivan Carvalho** - [IvanIsCoding](https://github.com/IvanIsCoding)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details