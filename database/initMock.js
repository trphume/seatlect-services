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
    email: 'fake1@email.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    dob: new Date('2000-10-15'),
    avatar: '',
    favorite: [],
    verified: false,
  },
  {
    _id: samuelId,
    username: 'Samuel',
    email: 'fake2@email.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    dob: new Date('1999-07-10'),
    avatar: '',
    favorite: [],
    verified: false,
  },
  {
    _id: gunId,
    username: 'Gun',
    email: 'fake3@email.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    dob: new Date('2004-02-22'),
    avatar: '',
    favorite: [],
    verified: false,
  },
]);


// business collection
db.business.insertMany([
  {
    _id: brightioId,
    username: 'Brightio',
    email: 'brightio@email.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'Brightio',
    type: 'Bar',
    tags: ['BAR', 'JAPANESE', 'LIVE MUSIC'],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [100.769652, 13.727892]
    },
    address: 'Keki Ngam 4, Chalong Krung 1, Latkrabang, Bangkok, 10520',
    displayImage: '',
    images: [],
    placement: [],
    menu: [],
    status: 1,
  },
  {
    _id: beerBurgerId,
    username: 'BeerBurger',
    email: 'beerburgero@email.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'Beer and Burger',
    type: 'Restaurant',
    tags: ['BEER', 'BURGER', 'LIVE MUSIC'],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [100.765001, 13.727830]
    },
    address: '611 Chalong Krung 1, Latkrabang, Bangkok, 10520',
    displayImage: '',
    images: [],
    placement: [],
    menu: [],
    status: 1,
  },
  {
    _id: ironBuffetId,
    username: 'IronBuffet',
    email: 'ironbuffet@email.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'Iron Buffet',
    type: 'Restaurant',
    tags: ['STEAK', 'BUFFET'],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [100.780103, 13.723117]
    },
    address: '44 Chalong Krung 1, Latkrabang, Bangkok 10520',
    displayImage: '',
    images: [],
    placement: [],
    menu: [],
    status: 1,
  },
  {
    _id: specialTaleId,
    username: 'SpecialTale',
    email: 'specialtale@email.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'SpecialTale',
    type: 'Bar',
    tags: ['COCKTAIL', 'BAR', 'LIVE MUSIC'],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [99.780103, 14.723117]
    },
    address: 'this is honestly, just some made up address',
    displayImage: '',
    images: [],
    placement: [],
    menu: [],
    status: 1
  }
]);