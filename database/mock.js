// All password in the mock api is 'ExamplePassword123'
// Uses Bcrypt with 12 round for encryption

// customer collection
db.customer.insertMany([
  {
    name: 'Jake',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    dob: new Date('2000-10-15'),
    avatar: '',
    preference: [],
    favorite: []
  },
  {
    name: 'Samuel',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    dob: new Date('1999-07-10'),
    avatar: '',
    preference: [],
    favorite: []
  },
  {
    name: 'Gun',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    dob: new Date('2004-02-22'),
    avatar: '',
    preference: [],
    favorite: []
  },
]);


// business collection
db.business.insertMany([
  {
    name: 'Brightio',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'Brightio',
    type: ['CONTEMPORARY', 'BAR', 'JAPANESE', 'FRENCH', 'LIVE MUSIC'],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.',
    location: {
      type: 'Point',
      coordinates: [13.727892, 100.769652]
    },
    address: 'Keki Ngam 4, Chalong Krung 1, Latkrabang, Bangkok, 10520',
    displayImage: '',
    images: [],
    placement: [],
    menu: [],
    menu_item: [],
    policy: []
  },
  {
    name: 'BeerBurger',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'Beer and Burger',
    type: ['PUB', 'FAST CASUAL', 'BURGER', 'LIVE MUSIC'],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.',
    location: {
      type: 'Point',
      coordinates: [13.727830, 100.765001]
    },
    address: '611 Chalong Krung 1, Latkrabang, Bangkok, 10520',
    displayImage: '',
    images: [],
    placement: [],
    menu: [],
    menu_item: [],
    policy: []
  },
  {
    name: 'IronBuffet',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'Iron Buffet',
    type: ['STEAK', 'BUFFET'],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.',
    location: {
      type: 'Point',
      coordinates: [13.723117, 100.780103]
    },
    address: '44 Chalong Krung 1, Latkrabang, Bangkok 10520',
    displayImage: '',
    images: [],
    placement: [],
    menu: [],
    menu_item: [],
    policy: []
  }
]);