# Stock Market Information Website

A modern, responsive website for stock market information and analysis, built with Go (Fiber) for the backend and vanilla HTML/CSS/JavaScript for the frontend.

## Features

- **Real-time Stock Data**: Display current stock prices, changes, and percentage movements
- **Watchlist Management**: Add and track your favorite stocks
- **Returns Calculator**: Calculate potential returns on investments with compound interest
- **Market Analysis**: Access detailed market insights and trends
- **IPO Calendar**: Stay updated with upcoming Initial Public Offerings
- **Futures & Options**: Track F&O market data
- **SIP/Mutual Funds**: Explore systematic investment plans and mutual fund options
- **Team Information**: Learn about our expert team members
- **Contact Form**: Get in touch with our support team

## Project Structure

```
Stocks/
├── main.go              # Go backend server
├── README.md           # Project documentation
└── src/                # Frontend files
    ├── index.html      # Main landing page
    ├── css/            # Stylesheets
    │   └── style.css   # Main styles
    ├── js/             # JavaScript files
    │   └── main.js     # Main JavaScript
    └── pages/          # Additional pages
        ├── about.html
        ├── contact.html
        ├── team.html
        └── login.html
```

## Tech Stack

- **Backend**: Go with Fiber framework
- **Frontend**: HTML5, CSS3, JavaScript
- **Styling**: Custom CSS with modern design principles
- **Icons**: Font Awesome
- **Fonts**: Google Fonts (Poppins)

## Setup Instructions

1. **Prerequisites**
   - Go 1.16 or higher
   - Git

2. **Installation**
   ```bash
   # Clone the repository
   git clone <repository-url>
   cd Stocks

   # Install Go dependencies
   go mod tidy

   # Run the server
   go run main.go
   ```

3. **Access the Website**
   - Open your browser and navigate to `http://localhost:3000`
   - The website will be served from the `src` directory

## API Endpoints

- `GET /api/stocks` - Get current stock data
- `GET /api/watchlist` - Get user's watchlist
- `POST /api/watchlist` - Add stock to watchlist
- `POST /api/calculate-returns` - Calculate investment returns

## Features in Detail

### 1. Stock Data Display
- Real-time stock prices
- Price changes and percentage movements
- Interactive stock cards with hover effects

### 2. Watchlist
- Add/remove stocks from watchlist
- Persistent storage of favorite stocks
- Quick access to important stock information

### 3. Returns Calculator
- Calculate future value of investments
- Consider initial investment and monthly contributions
- Account for compound interest
- Visual representation of potential returns

### 4. Market Analysis
- Technical analysis tools
- Market trends and patterns
- Historical data visualization

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

For any queries or support, please reach out through the contact form on the website or email us at support@stockmarket.com 