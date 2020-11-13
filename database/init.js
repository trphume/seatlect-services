db.createCollection('customer', {
  validator: {
    $jsonSchema: {
      bsonType: 'object', required: ['name', 'password', 'dob', 'avatar', 'preference', 'favorite'], properties: {
        name: {
          bsonType: 'string', minLength: 3,
          maxLength: 20
        }, password: {
          bsonType: 'string', minLength: 60,
          maxLength: 60
        }, dob: { bsonType: 'date' }, avatar: { bsonType: 'string' }, preference: {
          bsonType: 'array', items: {
            bsonType: 'string', uniqueItems: true,
            maxItems: 5
          }
        }, favorite: {
          bsonType: 'array', items: {
            bsonType: 'objectId', uniqueItems: true,
            maxItems: 5
          }
        }
      }
    }
  }
});

db.createCollection('business', {
  validator: {
    $jsonSchema: {
      bsonType: 'object', required: ['name', 'password', 'businessName', 'type', 'description', 'location', 'address', 'displayImage', 'images', 'placement', 'menu', 'policy'], properties: {
        name: {
          bsonType: 'string', minLength: 3,
          maxLength: 20
        }, password: {
          bsonType: 'string', minLength: 60,
          maxLength: 60
        }, businessName: {
          bsonType: 'string', minLength: 3,
          maxLength: 50
        }, type: {
          bsonType: 'array', items: {
            bsonType: 'string', uniqueItems: true,
            maxItems: 5
          }
        }, description: { bsonType: 'string', maxLength: 50 }, location: {
          bsonType: 'object',
          required: ['type', 'coordinates'], properties: {
            type: { bsonType: 'string', enum: ['Point'] }, coordinates: {
              bsonType: 'array', items: {
                bsonType: 'double', minItems: 2,
                maxItems: 2,
                items: [{ bsonType: 'double', minimum: -180, maximum: 180 }, { bsonType: 'double', minimum: -90, maximum: 90 }]
              }
            }
          }
        }, address: {
          bsonType: 'string', minLength: 3,
          maxLength: 50
        }, displayImage: { bsonType: 'string' }, images: {
          bsonType: 'array', items: {
            bsonType: 'string', uniqueItems: true,
            maxItems: 10
          }
        }, placement: {
          bsonType: 'array', maxItems: 5, items: {
            required: ['name', 'entity', 'default'], properties: {
              name: {
                bsonType: 'string', minLength: 3,
                maxLength: 20
              }, entity: {
                bsonType: 'array', items: {
                  required: ['name', 'floor', 'type', 'reserved'], properties: {
                    name: {
                      bsonType: 'string', minLength: 1,
                      maxLength: 2
                    }, floor: { bsonType: 'int', min: 1 }, type: { bsonType: 'string', enum: ['SEAT', 'TABLE', 'STAIRS', 'TOILET', 'STAGE', 'BLOCK'] }, reserved: { bsonType: 'bool' }, x: { bsonType: 'double' }, y: { bsonType: 'double' }
                  }
                }
              }, default: { bsonType: 'bool' }
            }
          }
        }, menu: {
          bsonType: 'array', maxItems: 5, items: {
            required: ['name', 'description', 'items', 'default'], properties: {
              name: {
                bsonType: 'string', minLength: 3,
                maxLength: 20
              }, description: { bsonType: 'string', maxLength: 125 }, items: {
                bsonType: 'array', uniqueItems: true,
                maxItems: 100, items: {
                  required: ['name', 'description', 'image', 'price'], properties: {
                    name: {
                      bsonType: 'string', minLength: 3,
                      maxLength: 20
                    }, description: { bsonType: 'string', maxLength: 125 }, image: { bsonType: 'string' }, price: { bsonType: 'decimal', min: 0 }
                  }
                }
              }, default: { bsonType: 'bool' }
            }
          }
        }, policy: {
          bsonType: 'array', maxItems: 100, items: {
            required: ['name', 'description', 'before', 'freeCancelDeadline', 'cancelRate', 'basePrice'], properties: {
              name: {
                bsonType: 'string', minLength: 3,
                maxLength: 20
              }, description: { bsonType: 'string', maxLength: 125 }, before: { bsonType: 'int', min: 0 }, freeCancelDeadline: { bsonType: 'int', min: 0, }, cancelRate: {
                bsonType: 'decimal', min: 0,
                max: 100
              }, basePrice: { bsonType: 'decimal', min: 0 }
            }
          }
        }
      }
    }
  }
});

db.createCollection('reservation', {
  validator: {
    $jsonSchema: {
      bsonType: 'object', required: ['businessId', 'name', 'start', 'end', 'placement', 'menu', 'policy'], properties: {
        businessId: { bsonType: 'objectId' }, name: {
          bsonType: 'string', minLength: 3,
          maxLength: 20
        }, start: { bsonType: 'date' }, end: { bsonType: 'date' }, placement: {
          bsonType: 'object',
          required: ['name', 'entity', 'default'], properties: {
            name: {
              bsonType: 'string', minLength: 3,
              maxLength: 20
            }, entity: {
              bsonType: 'array', items: {
                required: ['name', 'floor', 'type', 'reserved'], properties: {
                  name: {
                    bsonType: 'string', minLength: 1,
                    maxLength: 2
                  }, floor: { bsonType: 'int', min: 1 }, type: { bsonType: 'string', enum: ['SEAT', 'TABLE', 'STAIRS', 'TOILET', 'STAGE', 'BLOCK'] }, reserved: { bsonType: 'bool' }, x: { bsonType: 'double' }, y: { bsonType: 'double' }
                }
              }
            }, default: { bsonType: 'bool' }
          }
        }, menu: {
          bsonType: 'array', items: {
            required: ['name', 'description', 'image', 'price'], properties: {
              name: {
                bsonType: 'string', minLength: 3,
                maxLength: 20
              }, description: { bsonType: 'string', maxLength: 125 }, image: { bsonType: 'string' }, price: { bsonType: 'decimal', min: 0 }
            }
          }
        }, policy: {
          bsonType: 'object',
          required: ['name', 'description', 'before', 'freeCancelDeadline', 'cancelRate', 'basePrice'], properties: {
            name: {
              bsonType: 'string', minLength: 3,
              maxLength: 20
            }, description: { bsonType: 'string', maxLength: 125 }, before: { bsonType: 'int', min: 0 }, freeCancelDeadline: { bsonType: 'int', min: 0, }, cancelRate: {
              bsonType: 'decimal', min: 0,
              max: 100
            }, basePrice: { bsonType: 'decimal', min: 0 }
          }
        }
      }
    }
  }
});

db.createCollection('order', {
  validator: {
    $jsonSchema: {
      bsonType: 'object', required: ['customerId', 'businessId', 'paymentDate', 'start', 'end', 'reserve', 'preorder', 'basePrice', 'totalPrice', 'status'], properties: {
        customerId: { bsonType: 'objectId' }, businessId: { bsonType: 'objectId' }, paymentDate: { bsonType: 'date' }, start: { bsonType: 'date' }, end: { bsonType: 'date' }, reserve: { bsonType: 'array', items: { bsonType: 'string' } }, preorder: {
          bsonType: 'array', items: {
            required: ['name', 'quantity', 'price'], properties: {
              name: {
                bsonType: 'string', minLength: 3,
                maxLength: 20
              }, quantity: { bsonType: 'int', min: 1 }, price: { bsonType: 'decimal', min: 0 }
            }
          }
        }, basePrice: { bsonType: 'decimal', min: 0 }, totalPrice: { bsonType: 'decimal', min: 0 }, status: { bsonType: 'string', enum: ['PAID', 'USED', 'EXPIRED', 'CANCELLED'] }
      }
    }
  }
});  