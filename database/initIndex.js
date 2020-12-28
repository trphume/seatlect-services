db.customer.createIndex({ username: 1 }, { unique: true });

db.business.createIndex({ username: 1 }, { unique: true });
db.business.createIndex({ location: "2dsphere" });

db.reservation.createIndex({ businessId: 1 });

db.order.createIndex({ customerId: 1 });
db.order.createIndex({ businessId: 1 });