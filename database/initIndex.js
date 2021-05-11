// customer
db.customer.createIndex({ username: 1 }, { unique: true });
db.customer.createIndex({ email: 1 }, { unique: true });

// business
db.business.createIndex({ username: 1 }, { unique: true });
db.business.createIndex({ email: 1 }, { unique: true });
db.business.createIndex({businessName: "text"})
db.business.createIndex({ location: "2dsphere" });

// reservation
db.reservation.createIndex({ businessId: 1 });
db.reservation.createIndex({name: "text"})
db.reservation.createIndex({location: "2dsphere"});

// customer
db.order.createIndex({ customerId: 1 });
db.order.createIndex({ businessId: 1 });