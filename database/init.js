
// customer collection
db.createCollection('customer', { validator: { $jsonSchema: { bsonType: 'object', required: ['name', 'password', 'dob', 'avatar', 'preference', 'favorite'], properties: { name: { bsonType: 'string' }, password: { bsonType: 'string' }, dob: { bsonType: 'date' }, avatar: { bsonType: 'string' }, preference: { bsonType: 'array', items: { bsonType: 'string' } }, favorite: { bsonType: 'array', items: { bsonType: 'objectId' } } } } } });

db.customer.createIndex({ name: 1 }, { unique: true });

// business collection
db.createCollection('business', {
  validator: {
    $jsonSchema: {
      bsonType: 'object', required: ['name', 'password', 'businessName', 'type', 'description', 'location', 'address', 'displayImage', 'images', 'placement', 'menu', 'menu_item', 'policy'], properties: {
        name: { bsonType: 'string' }, password: { bsonType: 'string' }, businessName: { bsonType: 'string' }, type: { bsonType: 'string' }, description: { bsonType: 'string' }, location: {
          bsonType: 'object',
          required: ['type', 'coordinates'], properties: { type: { bsonType: 'string' }, coordinates: { bsonType: 'array', items: { bsonType: 'double' } } }
        }, address: { bsonType: 'string' }, displayImage: { bsonType: 'string' }, images: { bsonType: 'array', items: { bsonType: 'string' } }, placement: {
          bsonType: 'array', items: {
            required: ['name', 'entity', 'default'], properties: {
              name: { bsonType: 'string' }, entity: {
                bsonType: 'array', items: {
                  required: ['id', 'floor', 'type', 'reserved'], properties: { id: { bsonType: 'string' }, floor: { bsonType: 'int' }, type: { bsonType: 'string' }, reserved: { bsonType: 'bool' } }
                }
              }, default: { bsonType: 'bool' }
            }
          }
        }, menu: {
          bsonType: 'array', items: {
            required: ['name', 'description', 'items', 'default'], properties: { name: { bsonType: 'string' }, description: { bsonType: 'string' }, items: { bsonType: 'array', items: { bsonType: 'string' } }, default: { bsonType: 'bool' } }
          }
        }, menu_item: {
          bsonType: 'array', items: {
            required: ['name', 'description', 'image', 'price'], properties: { name: { bsonType: 'string' }, description: { bsonType: 'string' }, image: { bsonType: 'string' }, price: { bsonType: 'decimal' } }
          }
        }, policy: {
          bsonType: 'array', items: {
            required: ['name', 'description', 'before', 'freeCancelDeadline', 'cancelRate', 'basePrice'], properties: { name: { bsonType: 'string' }, description: { bsonType: 'string' }, before: { bsonType: 'int' }, freeCancelDeadline: { bsonType: 'int' }, cancelRate: { bsonType: 'decimal' }, basePrice: { bsonType: 'decimal' } }
          }
        }
      }
    }
  }
});

db.business.createIndex({ name: 1 }, { unique: true });

// reservation collection
db.createCollection('reservation', {
  validator: {
    $jsonSchema: {
      bsonType: 'object', required: ['businessId', 'name', 'start', 'end', 'placement', 'menu_item', 'policy'], properties: {
        businessId: { bsonType: 'objectId' }, name: { bsonType: 'string' }, start: { bsonType: 'date' }, end: { bsonType: 'date' }, placement: {
          bsonType: 'object',
          required: ['name', 'entity', 'default'], properties: {
            name: { bsonType: 'string' }, entity: {
              bsonType: 'array', items: {
                required: ['id', 'floor', 'type', 'reserved'], properties: { id: { bsonType: 'string' }, floor: { bsonType: 'int' }, type: { bsonType: 'string' }, reserved: { bsonType: 'bool' } }
              }
            }, default: { bsonType: 'bool' }
          }
        }, menu_item: {
          bsonType: 'array', items: {
            required: ['name', 'description', 'image', 'price'], properties: { name: { bsonType: 'string' }, description: { bsonType: 'string' }, image: { bsonType: 'string' }, price: { bsonType: 'decimal' } }
          }
        }, policy: {
          bsonType: 'object',
          required: ['name', 'description', 'before', 'freeCancelDeadline', 'cancelRate', 'basePrice'], properties: { name: { bsonType: 'string' }, description: { bsonType: 'string' }, before: { bsonType: 'int' }, freeCancelDeadline: { bsonType: 'int' }, cancelRate: { bsonType: 'decimal' }, basePrice: { bsonType: 'decimal' } }
        }
      }
    }
  }
});

db.reservation.createIndex({ businessId: 1 });

// order collection
db.createCollection('order', {
  validator: {
    $jsonSchema: {
      bsonType: 'object', required: ['customerId', 'businessId', 'paymentDate', 'start', 'end', 'reserve', 'item', 'basePrice', 'totalPrice', 'status'], properties: {
        customerId: { bsonType: 'objectId' }, businessId: { bsonType: 'objectId' }, paymentDate: { bsonType: 'date' }, start: { bsonType: 'date' }, end: { bsonType: 'date' }, reserve: { bsonType: 'array', items: { bsonType: 'string' } }, item: {
          bsonType: 'array', items: {
            required: ['name', 'description', 'image', 'price'], properties: { name: { bsonType: 'string' }, description: { bsonType: 'string' }, image: { bsonType: 'string' }, price: { bsonType: 'decimal' } }
          }
        }, basePrice: { bsonType: 'decimal' }, totalPrice: { bsonType: 'decimal' }, status: { bsonType: 'string' }
      }
    }
  }
});

db.order.createIndex({ customerId: 1 });
db.order.createIndex({ businessId: 1 });