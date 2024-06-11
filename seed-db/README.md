# Seeding local dev mongo with initial data

## Description

This project is a simple script to seed a database with data. 
It supports a `--forced-seed` flag to force seeding, which can also be set through the `FORCED_SEED` environment variable.
By using force seed the script will empty the db first to apply the default data.

## Installation

To install the necessary dependencies, run the following command:

```make install```

## Usage
To seed the database, run the following command:

```make run```

To force seeding, you can either set the FORCED_SEED environment variable to "true" or include the --forced-seed flag when running the script:

```make run-forced-seed```

## Cleaning Up
To remove the node_modules directory, run the following command:

```make clean```

## Contributing
If you want to contribute to this project, please send a pull request.

## License
This project is licensed under the MIT License.