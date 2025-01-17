# go-shopping-list

## Summary

Do you find making shopping lists BORING? Because I sure do... This repo is to automate the creation of shopping lists and to practice my Golang skills! 

## Demo

When you are running on mac and have the shopping.workflow file present this is how the code operates:

https://user-images.githubusercontent.com/32711718/189308917-77185ca7-6811-4a2f-b3c4-a85158703dde.mov

## How does it work? 

1. The codebase works by firstly reading all the JSON files in the recipes folder. 
2. Once the JSON's have been processed they are passed to the GUI and displayed in a list
3. When one of the list items is clicked it will run automation scripts which add a reminder to my reminders app
4. The reminders app is an app available on mac and iphone.
5. The iCloud then syncs the items in my list to my phone so I can go to the shops and get the stuff I need.

## Maintaining Professional Standards

To ensure my code is professional and extendable I followed these rules when making changes:

1. Apply unit testing wherever possible and aim for 80% coverage in packages
2. Use `golangci-lint run` and the revive linter with all rules enabled: https://github.com/mgechev/revive 
3. Using Interfaces to mock results using the following module: https://github.com/vektra/mockery 

## What happens if I'm not running on a Mac?

If you are not running on a macbook and don't have the shopping.workflow present then the system will just print that it is pretending to add the ingredients so you can still use the system. See the terminal demo below:

https://user-images.githubusercontent.com/32711718/189309910-072d7b0d-bffa-4661-ad5a-01fb5aaff30e.mov

