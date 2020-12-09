// All password in the mock api is 'ExamplePassword123'
// Uses Bcrypt with 12 round for encryption


// Ids
jakeId = ObjectId('5facaf3bd646b77f40481343');
samuelId = ObjectId('5facaf7a35c1e1db56597485');
gunId = ObjectId('5facaf818b4f49b3cf1f1792');

brightioId = ObjectId('5facafef6b28446f285d7ae4');
beerBurgerId = ObjectId('5facaff31c6d49b2c7256bf3');
ironBuffetId = ObjectId('5facaff9e4d46967c9c2a558');
specialTaleId = ObjectId('5fcde2ec209efa45620a08b6');

// customer collection
db.customer.insertMany([
  {
    _id: jakeId,
    username: 'Jake',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    dob: new Date('2000-10-15'),
    avatar: '',
    preference: [],
    favorite: []
  },
  {
    _id: samuelId,
    username: 'Samuel',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    dob: new Date('1999-07-10'),
    avatar: '',
    preference: [],
    favorite: []
  },
  {
    _id: gunId,
    username: 'Gun',
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
    _id: brightioId,
    username: 'Brightio',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'Brightio',
    tags: [BAR', 'JAPANESE', 'LIVE MUSIC'],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
      location: {
        type: 'Point',
        coordinates: [13.727892, 100.769652]
      },
      address: 'Keki Ngam 4, Chalong Krung 1, Latkrabang, Bangkok, 10520',
      displayImage: '',
      images: [],
      placement: [],
      menu: [],
      policy: { minAge: 0 }
  },
  {
    _id: beerBurgerId,
    username: 'BeerBurger',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'Beer and Burger',
    tags: ['BEER', 'BURGER', 'LIVE MUSIC'],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [13.727830, 100.765001]
    },
    address: '611 Chalong Krung 1, Latkrabang, Bangkok, 10520',
    displayImage: '',
    images: [],
    placement: [],
    menu: [],
    policy: { minAge: 0 }
  },
  {
    _id: ironBuffetId,
    username: 'IronBuffet',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'Iron Buffet',
    tags: ['STEAK', 'BUFFET'],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [13.723117, 100.780103]
    },
    address: '44 Chalong Krung 1, Latkrabang, Bangkok 10520',
    displayImage: '',
    images: [],
    placement: [],
    menu: [],
    policy: { minAge: 0 }
  },
  {
    _id: specialTaleId,
    username: 'SpecialTale',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'SpecialTale',
    tags: ['COCKTAIL', 'BAR', 'LIVE MUSIC'],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [14.723117, 99.780103]
    },
    address: 'this is honestly, just some made up address',
    displayImage: '',
    images: [],
    placement: [],
    menu: [],
    policy: { minAge: 21 }
  }
]);