# BeMyRoomie - Find Your Ideal Roommate & Rental

BeMyRoomie is a web-based platform that simplifies the process of finding a roommate or rental accommodation. It provides users with an easy way to list, search, and filter available flats, rooms, or roommates based on preferences and requirements.

## Problem Statement
Finding a suitable roommate or rental property can be a challenging task. Many existing platforms like Facebook and MagicBricks lack advanced filtering options, involve extensive research, and pose risks of scams. Communication can also be cumbersome. **BeMyRoomie** aims to solve these issues by centralizing the process with smart filters and seamless interactions.

## Features
- **Find a Roommate**: Search for a compatible roommate.
- **Find a Flat**: Rent a flat as an individual or with a group.
- **List a Property**: Post vacant rooms or flats for rent.
- **Search & Filter**: Find accommodations based on location, gender, budget, and preferences.
- **Interest Requests**: Express interest in listings and get accepted/rejected by listers.
- **User Profiles**: Maintain preferences and details like smoking habits, dietary preferences, and lifestyle.

## Tech Stack
- **Backend**: Golang
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Token)

## Prerequisites
Ensure you have the following installed:
- **Go** (Download: https://golang.org/dl/)
- **PostgreSQL** (Download: https://www.postgresql.org/download/)

Install dependencies using:
```
 go mod tidy
```

## Project Setup
1. Clone the repository:
   ```sh
   git clone <repository_url>
   cd bemyroomie
   ```
2. Create a `config.go` and add HTTP:PORT and JWT_SECRET_KEY
3. Run the project:
   `cd cmd`
   `go run main.go` 

## API Endpoints
### Authentication
- **Lister Signup**: `POST /lister/signup`
- **Finder Signup**: `POST /finder/signup`
- **User Sign-in**: `POST /signin`

### Profile
- **Update Profile**: `PATCH /profile`
- **Set Preferences**: `POST /preferences`

### Listings
- **View Listings**: `GET /listings`
- **Get Specific Listing**: `GET /listings/{listing_id}`
- **Post a Listing**: `POST /listing`
- **Update Listing**: `PUT /listings/{listing_id}`

### Interests
- **Show Interest in Listing**: `POST /interest/{property_id}`
- **View Interested Listings**: `GET /interests/listings`
- **Accept/Reject Interest**: `PATCH /interests/listings/{listing_id}`

## Notes
- A user (lister) can have only one active listing at a time.
- Proper authentication and authorization are enforced using JWT.
- More filtering features will be added in future updates.

## Contributing
Feel free to contribute by submitting issues or pull requests.

## License
This project is licensed under the MIT License.



