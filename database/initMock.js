// All password in the mock api is 'ExamplePassword'
// Uses Bcrypt with 12 round for encryption


// Ids
admin1 = ObjectId('604dfa455226a8714411f33d');

jakeId = ObjectId('5facaf3bd646b77f40481343');
samuelId = ObjectId('5facaf7a35c1e1db56597485');
gunId = ObjectId('5facaf818b4f49b3cf1f1792');

brightioId = ObjectId('5facafef6b28446f285d7ae4');
centralBrightioId = ObjectId('6083f7e4630ee2a709fc8234');
jiaShinId = ObjectId('608458b0a704b63fed6c7731')
beerBurgerId = ObjectId('5facaff31c6d49b2c7256bf3');
ironBuffetId = ObjectId('5facaff9e4d46967c9c2a558');
specialTaleId = ObjectId('5fcde2ec209efa45620a08b6');
reggaeBarId = ObjectId('60ae0864220ff22a378bd2f0');
anotherBrightioId = ObjectId('60b759196d645ef1f6a4ebc8');
tateCafeId = ObjectId('60b759a06b4ced4a5e1fa652');

reservation_1 = ObjectId('6035f3a48d505df0b9d043a3');
reservation_2 = ObjectId('604c80551714a597557abc2e');
reservation_3 = ObjectId('608f5977ce9b402eda2432be');
reservation_4 = ObjectId('608f59c053a970a744ce14ac');
reservation_5 = ObjectId('608f59c96f17b50ec1ad94e2');
reservation_6 = ObjectId('608f59d0e58318c67666fced');

order_1 = ObjectId('609a72efe0a21a5e7f6b758e');

orderA = ObjectId('6035fb35bf78e591bea86350');

// Reusable mockup
exampleDisplayImage = 'https://i.imgur.com/rXjqn0y.jpeg'

exampleImages = ['https://i.imgur.com/g17EY2i.jpg', 'https://i.imgur.com/RjFgQSZ.jpeg']

exampleMenu = [
  { name: "Fries", description: "Just fries", image: "https://i.imgur.com/rXjqn0y.jpeg", price: 10 },
  { name: "Salty Fries", description: "Just salty fries", image: "https://i.imgur.com/rXjqn0y.jpeg", price: 10 }
]

examplePlacement = {
  width: 800,
  height: 800,
  seats: [
    {
      name: 'Stage',
      floor: 1,
      type: 'OBJECT_SQUARE',
      space: 4,
      x: 10,
      y: 10,
      width: 80,
      height: 320,
      rotation: 0,
    },
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
}

reservation3Placement = {
  width: 800,
  height: 800,
  seats: [
    {
      name: 'Stage',
      floor: 1,
      type: 'OBJECT_SQUARE',
      space: 0,
      user: null,
      username: '',
      status: '',
      x: 10,
      y: 10,
      width: 80,
      height: 320,
      rotation: 0,
    },
    {
      name: 'A1',
      floor: 1,
      type: 'TABLE',
      space: 4,
      user: jakeId,
      username: 'Jake',
      status: 'TAKEN',
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
      user: null,
      username: '',
      status: 'AVAILABLE',
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
      user: null,
      username: '',
      status: 'AVAILABLE',
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
      user: null,
      username: '',
      status: 'AVAILABLE',
      x: 200,
      y: 200,
      width: 80,
      height: 80,
      rotation: 0,
    },
  ]
}

exampleEmployees = [
  { username: "EmployeeA", password: "ExamplePassword" },
  { username: "EmployeeB", password: "ExamplePassword" },
]

exampleReservationPlacement = {
  width: 800,
  height: 800,
  seats: [
    {
      name: 'Stage',
      floor: 1,
      type: 'OBJECT_SQUARE',
      space: 0,
      user: null,
      username: '',
      status: '',
      x: 10,
      y: 10,
      width: 80,
      height: 320,
      rotation: 0,
    },
    {
      name: 'A1',
      floor: 1,
      type: 'TABLE',
      space: 4,
      user: null,
      username: '',
      status: 'AVAILABLE',
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
      user: null,
      username: '',
      status: 'AVAILABLE',
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
      user: null,
      username: '',
      status: 'AVAILABLE',
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
      user: null,
      username: '',
      status: 'AVAILABLE',
      x: 200,
      y: 200,
      width: 80,
      height: 80,
      rotation: 0,
    },
  ]
}

emptyPlacement = {
  width: 800,
  height: 800,
  seats: []
}

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
    email: 'traphumem@gmail.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    dob: new Date('2000-10-15'),
    favorite: [],
    verified: false,
  },
  {
    _id: samuelId,
    username: 'Samuel',
    email: '60090043@kmitl.ac.th',
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
    tags: [],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [100.769652, 13.727892]
    },
    address: 'Keki Ngam 4, Chalong Krung 1, Latkrabang, Bangkok, 10520',
    displayImage: exampleDisplayImage,
    images: exampleImages,
    placement: examplePlacement,
    menu: exampleMenu,
    status: 1,
    employee: exampleEmployees,
  },
  {
    _id: beerBurgerId,
    username: 'BeerBurger',
    email: 'beerburgero@email.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'Beer and Burger',
    type: 'Restaurant',
    tags: [],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [100.765001, 13.727830]
    },
    address: '611 Chalong Krung 1, Latkrabang, Bangkok, 10520',
    displayImage: exampleDisplayImage,
    images: exampleImages,
    placement: examplePlacement,
    menu: exampleMenu,
    status: 1,
    employee: exampleEmployees,
  },
  {
    _id: ironBuffetId,
    username: 'IronBuffet',
    email: 'ironbuffet@email.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'Iron Buffet',
    type: 'Restaurant',
    tags: [],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [100.780103, 13.723117]
    },
    address: '44 Chalong Krung 1, Latkrabang, Bangkok 10520',
    displayImage: exampleDisplayImage,
    images: exampleImages,
    placement: examplePlacement,
    menu: exampleMenu,
    status: 1,
    employee: exampleEmployees,
  },
  {
    _id: specialTaleId,
    username: 'SpecialTale',
    email: 'specialtale@email.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'SpecialTale',
    type: 'Bar',
    tags: [],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [99.780103, 14.723117]
    },
    address: 'this is honestly, just some made up address',
    displayImage: exampleDisplayImage,
    images: exampleImages,
    placement: examplePlacement,
    menu: exampleMenu,
    status: 1,
    employee: exampleEmployees,
  },
  {
    _id: jiaShinId,
    username: 'JiaShin',
    email: 'jiashin@gmail.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'JiaShin',
    type: 'Bar',
    tags: [],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [100.5018, 13.7563]
    },
    address: 'Bangkok',
    displayImage: exampleDisplayImage,
    images: exampleImages,
    placement: examplePlacement,
    menu: exampleMenu,
    status: 1,
    employee: exampleEmployees,
  },
  {
    _id: centralBrightioId,
    username: 'CentralBrightio',
    email: 'traphumedev@gmail.com',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'Central Brightio',
    type: 'Bar',
    tags: [],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [100.53793107547114, 13.745226384751511]
    },
    address: 'Groove, Central World',
    displayImage: exampleDisplayImage,
    images: exampleImages,
    placement: examplePlacement,
    menu: exampleMenu,
    status: 1,
    employee: exampleEmployees,
  },
  {
    _id: reggaeBarId,
    username: 'ReggaeBar',
    email: '60090043@kmitl.ac.th',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'ReggaeBar',
    type: 'Bar',
    tags: [],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [100.5018, 13.7563]
    },
    address: 'Bangkok',
    displayImage: "",
    images: [],
    placement: emptyPlacement,
    menu: [],
    status: 0,
    employee: [],
  },
  {
    _id: anotherBrightioId,
    username: 'AnotherBrightio',
    email: '60090027@kmitl.ac.th',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'AnotherBrightio',
    type: 'Restaurant',
    tags: [],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [100.769452, 13.725892]
    },
    address: 'Keki Ngam 4, Chalong Krung 1, Latkrabang, Bangkok, 10520',
    displayImage: "",
    images: [],
    placement: emptyPlacement,
    menu: [],
    status: 0,
    employee: [],
  },
  {
    _id: tateCafeId,
    username: 'TateCafe',
    email: '60090011@kmitl.ac.th',
    password: '$2y$12$dx/ILJHQbxtQHDq04JAk/OICg25Cj9DmYv33FgYXfDa4gxOwJVJ9.',
    businessName: 'TateCafe',
    type: 'Restaurant',
    tags: [],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [100.5018, 13.7563]
    },
    address: 'Bangkok',
    displayImage: "",
    images: [],
    placement: emptyPlacement,
    menu: [],
    status: 0,
    employee: [],
  },
]);

// reservation
db.reservation.insertMany([
  {
    _id: reservation_1,
    businessId: brightioId,
    name: 'Brightio',
    start: new Date('2021-06-16T19:00:00Z'),
    end: new Date('2021-06-16T00:23:00Z'),
    placement: exampleReservationPlacement,
    image: exampleDisplayImage,
    location: {
      type: 'Point',
      coordinates: [100.769652, 13.727892]
    },
    type: 'Bar',
    status: 1,
  },
  {
    _id: reservation_2,
    businessId: jiaShinId,
    name: 'JiaShin',
    start: new Date('2021-06-16T19:00:00Z'),
    end: new Date('2021-06-16T23:00:00Z'),
    placement: exampleReservationPlacement,
    image: exampleDisplayImage,
    location: {
      type: 'Point',
      coordinates: [100.5018, 13.7563]
    },
    type: 'Bar',
    status: 1,
  },
  {
    _id: reservation_3,
    businessId: centralBrightioId,
    name: 'Central Brightio',
    start: new Date('2021-06-16T19:00:00Z'),
    end: new Date('2021-06-16T23:00:00Z'),
    placement: reservation3Placement,
    image: exampleDisplayImage,
    location: {
      type: 'Point',
      coordinates: [100.53793107547114, 13.745226384751511]
    },
    type: 'Bar',
    status: 1,
  },
  {
    _id: reservation_4,
    businessId: ironBuffetId,
    name: 'Iron Buffet',
    start: new Date('2021-06-16T19:00:00Z'),
    end: new Date('2021-06-16T23:00:00Z'),
    placement: exampleReservationPlacement,
    image: exampleDisplayImage,
    location: {
      type: 'Point',
      coordinates: [100.780103, 13.723117]
    },
    type: 'Restaurant',
    status: 1,
  },
  {
    _id: reservation_5,
    businessId: specialTaleId,
    name: 'Special Tale',
    start: new Date('2021-06-16T19:00:00Z'),
    end: new Date('2021-06-16T23:00:00Z'),
    placement: exampleReservationPlacement,
    image: exampleDisplayImage,
    location: {
      type: 'Point',
      coordinates: [99.780103, 14.723117]
    },
    type: 'Bar',
    status: 1,
  },
  {
    _id: reservation_6,
    businessId: beerBurgerId,
    name: 'Beer Burger',
    start: new Date('2021-06-16T19:00:00Z'),
    end: new Date('2021-06-16T23:00:00Z'),
    placement: exampleReservationPlacement,
    image: exampleDisplayImage,
    location: {
      type: 'Point',
      coordinates: [100.765001, 13.727830]
    },
    type: 'Restaurant',
    status: 1,
  },
]);

// order
db.order.insertMany([
  {
    _id: order_1,
    reservationId: reservation_3,
    customerId: jakeId,
    businessId: centralBrightioId,
    start: new Date('2021-06-16T19:00:00Z'),
    end: new Date('2021-06-16T23:00:00Z'),
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
      }
    ],
    status: 'AVAILABLE',
    image: 'https://i.imgur.com/rXjqn0y.jpeg',
    extraSpace: 0,
    name: 'Central Brightio'
  }
]);

// request
db.request.insertMany([
  {
    _id: brightioId,
    businessName: 'Brightio',
    type: 'Restaurant',
    tags: [],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [100.769652, 13.727892]
    },
    address: 'Keki Ngam 4, Chalong Krung 1, Latkrabang, Bangkok, 10520',
    createdAt: new Date('2021-04-25T19:00:00Z')
  },
  {
    _id: centralBrightioId,
    businessName: 'Super Central Brightio',
    type: 'Restaurant',
    tags: [],
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    location: {
      type: 'Point',
      coordinates: [100.53793107547114, 13.745226384751511]
    },
    address: 'Groove, Central World',
    createdAt: new Date('2021-04-25T19:00:00Z')
  }
]);