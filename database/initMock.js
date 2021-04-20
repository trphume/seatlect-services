// All password in the mock api is 'ExamplePassword'
// Uses Bcrypt with 12 round for encryption


// Ids
admin1 = ObjectId('604dfa455226a8714411f33d');

jakeId = ObjectId('5facaf3bd646b77f40481343');
samuelId = ObjectId('5facaf7a35c1e1db56597485');
gunId = ObjectId('5facaf818b4f49b3cf1f1792');

brightioId = ObjectId('5facafef6b28446f285d7ae4');
beerBurgerId = ObjectId('5facaff31c6d49b2c7256bf3');
ironBuffetId = ObjectId('5facaff9e4d46967c9c2a558');
specialTaleId = ObjectId('5fcde2ec209efa45620a08b6');

reservationA = ObjectId('6035f3a48d505df0b9d043a3');
reservationB = ObjectId('604c80551714a597557abc2e');

orderA = ObjectId('6035fb35bf78e591bea86350');

// admin collection
db.admin.insertMany([
  {
    _id: admin1,
    username: 'admin1',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
  }
]);

// customer collection
db.customer.insertMany([
  {
    _id: jakeId,
    username: 'Jake',
    email: 'fake1@email.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    dob: new Date('2000-10-15'),
    favorite: [],
    verified: false,
  },
  {
    _id: samuelId,
    username: 'Samuel',
    email: 'fake2@email.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    dob: new Date('1999-07-10'),
    favorite: [],
    verified: false,
  },
  {
    _id: gunId,
    username: 'Gun',
    email: 'fake3@email.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    dob: new Date('2004-02-22'),
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
    placement: {
      width: 800,
      height: 800,
      seats: []
    },
    menu: [],
    status: 1,
    employee: [],
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
    placement: {
      width: 800,
      height: 800,
      seats: []
    },
    menu: [
      { name: "Fries", description: "Just fries", image: "", price: 10 },
      { name: "Salty Fries", description: "Just salty fries", image: "", price: 10 }
    ],
    status: 1,
    employee: [],
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
    placement: {
      width: 800,
      height: 800,
      seats: []
    },
    menu: [
      { name: "Fries", description: "Just fries", image: "", price: 10 },
      { name: "Salty Fries", description: "Just salty fries", image: "", price: 10 }
    ],
    status: 1,
    employee: [],
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
    placement: {
      width: 800,
      height: 800,
      seats: [
        {
          name: 'A1',
          floor: 1,
          type: 'TABLE',
          space: 4,
          x: 100,
          y: 100,
          width: 80,
          height: 80,
          rotation: 0,
        },
        {
          name: 'A2',
          floor: 1,
          type: 'TABLE',
          space: 4,
          x: 200,
          y: 100,
          width: 80,
          height: 80,
          rotation: 0,
        },
        {
          name: 'B1',
          floor: 1,
          type: 'TABLE',
          space: 4,
          x: 100,
          y: 200,
          width: 80,
          height: 80,
          rotation: 0,
        },
        {
          name: 'B2',
          floor: 1,
          type: 'TABLE',
          space: 4,
          x: 200,
          y: 200,
          width: 80,
          height: 80,
          rotation: 0,
        },
      ]
    },
    menu: [],
    status: 1,
    employee: [],
  }
]);

// reservation
db.reservation.insertMany([
  {
    _id: reservationA,
    businessId: brightioId,
    name: 'Brightio',
    start: new Date('2021-02-24T19:00:00Z'),
    end: new Date('2021-02-25T00:00:00Z'),
    placement: {
      width: 1000,
      height: 1000,
      seats: [
        {
          name: 'A1',
          floor: 1,
          type: 'TABLE',
          space: 4,
          user: jakeId,
          status: 'TAKEN',
          x: 471.1235,
          y: 124.2363,
          width: 80,
          height: 80,
          rotation: 0,
        },
        {
          name: 'A2',
          floor: 1,
          type: 'TABLE',
          space: 4,
          user: null,
          status: 'AVAILABLE',
          x: 470.2353,
          y: 368.8553,
          width: 80,
          height: 80,
          rotation: 0,
        },
      ]
    },
    image: '',
    location: {
      type: 'Point',
      coordinates: [100.769652, 13.727892]
    },
    type: 'Bar',
  },
  {
    _id: reservationB,
    businessId: specialTaleId,
    name: 'SpecialTale',
    start: new Date('2021-02-24T19:00:00Z'),
    end: new Date('2021-02-25T00:00:00Z'),
    placement: {
      width: 1000,
      height: 1000,
      seats: [
        {
          name: 'A1',
          floor: 1,
          type: 'TABLE',
          space: 4,
          user: null,
          status: 'AVAILABLE',
          x: 471.1235,
          y: 124.2363,
          width: 80,
          height: 80,
          rotation: 0,
        },
        {
          name: 'A2',
          floor: 1,
          type: 'TABLE',
          space: 4,
          user: null,
          status: 'AVAILABLE',
          x: 470.2353,
          y: 368.8553,
          width: 80,
          height: 80,
          rotation: 0,
        },
      ]
    },
    image: '',
    location: {
      type: 'Point',
      coordinates: [99.780103, 14.723117]
    },
    type: 'Bar',
  }
]);

// order
db.order.insertMany([
  {
    _id: orderA,
    reservationId: reservationA,
    customerId: jakeId,
    businessId: brightioId,
    start: new Date('2021-02-24T19:00:00Z'),
    end: new Date('2021-02-25T00:00:00Z'),
    seats: [
      {
        name: 'A1',
        floor: 1,
        type: 'TABLE',
        space: 4,
        x: 471.1235,
        y: 124.2363,
        width: 80,
        height: 80,
        rotation: 0,
      }
    ],
    status: 'USED',
    image: '',
    extraSpace: 0,
    name: 'Brightio'
  }
]);

// request
db.request.insertMany([
  {
    _id: brightioId,
    businessName: 'Brightio',
    type: 'Cool Bar',
    tags: ['BAR', 'JAPANESE', 'LIVE MUSIC'],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [100.769652, 13.727892]
    },
    address: 'Keki Ngam 4, Chalong Krung 1, Latkrabang, Bangkok, 10520',
    createdAt: new Date('2021-02-25T19:00:00Z')
  }
]);