# Database

- [Database](#database)
  - [**Overview**](#overview)
  - [**MongoDB Schema**](#mongodb-schema)
    - [`customer`](#customer)
    - [`business`](#business)
      - [`placement`](#placement)
      - [`menu`](#menu)
      - [`menu_item`](#menu_item)
      - [`policy`](#policy)
    - [`reservation`](#reservation)
    - [`order`](#order)
  - [**Scripts**](#scripts)

## **Overview**

This repository contains files related to creating, testing, deploying and maintaining Seatlect's databases. Seatlect uses MongoDB as its primary data store for structured data.

The convention of this section is that the schema is defined in a key-value where the field name is the key and the value is the data type. Index are identified by the square brackets []. Composite index are [COMPOSITE NUMBER] where the number is an integer indicating the composite index group.

## **MongoDB Schema**

This section contains the schema for MongoDB database. Each one corresponds to a single collection.

### `customer`

The **customer** collection contains information on *general users*, these are users which does own a business. They are the customer of a business. The following defines the schema for the document in the collection.

```json
{
  "_id": "ObjectId [UNIQUE]",
  "name": "String [UNIQUE]",
  "password": "String",
  "dob": "Date",
  "avatar": "String",
  "preference": "Array<String>",
  "favorite": "Array<ObjectId>"
}
```

- **_id** - This is MongoDB default uniquely generated id
- **name** - The name of the user, is used on authentication
- **password** - Hashed password in string format
- **dob** - Date of birth of the user
- **avatar** - Link to the avatar image asset of the user
- **preferences** - List of tags of a business type
- **favorites** - List of ids associated with a business

### `business`

The **business** collection contains information on *business users* and there business, these are users that owns a business and can configure their business store through our web application. The following defines the schema for the document in the collection.

```json
{
  "_id": "ObjectId [UNIQUE]",
  "name": "String [UNIQUE]",
  "password": "String",
  "businessName": "String",
  "type": "Array<String>",
  "description": "String",
  "location": {
    "type": "<GeoJSON Point>",
    "coordinates": "<coordinates>"
  },
  "address": "String",
  "displayImage": "String",
  "images": "Array<String>",
  "placement": "Array<placement>",
  "menu": "Array<menu>",
  "menu_item": "Array<menu_item>",
  "policy": "Array<policy>"
}
```

- **_id** - This is MongoDB default uniquely generated id
- **name** - The name of the user, is used on authentication
- **password** - Hashed password in string format
- **businessName** - The name of the business, is not unique
- **type** - The array of types associated with this business
- **description** - Short description of the business, will be displayed on the mobile application
- **location** - Mongo GeoJSON object, requires 2sphere index
- **address** - Address name of the business
- **displayImage** - Equivalent to a profile photo
- **images** - List of image resource url
- **placement** - Array of placement document objects
- **menu** -Array of menu document objects
- **menu_item** - Array of menu_item document
- **policy** - Array of policy document

Note that placement, menu and policy are **templates** which is used during the **creation of reservation document**. Templates are used as a base for information to be contained in a reservation. They can then be **modified** before finalization of a reservation object creation.

#### `placement`

The **placement** document contains information on the *floor layout template* of a business. Businesses can have several placement that they can choose to apply to their reservation schedule. The following define the schema for a placement document.

```json
{
  "name": "String",
  "entity": "Array<entity>",
  "default": "Boolean"
}
```

Below is the **entity** document used in placement.

```json
{
  "id": "String",
  "floor": "32-bit Integer",
  "type": "String",
  "reserved": "Boolean"
}
```

- **name** - This uniquely identifies a placement template within the array of other placement template in the business object
- **entity** - An array of entity describing the placement layout
  - **id** - This has to be unique and is set upon creation of entity
  - **floor** - Indicates which floor the entity should be placed on
  - **type** - This indicates how the system should interpret the entity, if type is RESERVE then the system interprets this as a reservable entity
  - **reserved** - If true then the entity has already been reserved by a customer
- **default** - If true indicates that this placement should be on display in the mobile application

#### `menu`

The **menu** document contains a *list of menu_item names* of a business. Businesses can have several menu that they choose to apply to their reservation schedule. The following define the schema for a placement document.

```json
{
  "name": "String",
  "description": "String",
  "items": "Array<String>",
  "default": "Boolean"
}
```

- **name** - This uniquely identifies a menu template within the array of other menu in the business object
- **description** - Short description of the menu template
- **items** - Array of menu_item name which can be found in menu_items array in the business object
- **default** - If true indicates that this menu should be on display in the mobile application

#### `menu_item`

The **menu_item** document contains information on a particular item.

```json
{
  "name": "String",
  "description": "String",
  "image": "String",
  "price": "Decimal"
}
```

- **name** - This uniquely identifies an item within the array of other menu_items in the business object
- **description** - Short description of the item
- **image** - Image resource url
- **price** - The price of a single order of this item

#### `policy`

The **policy** object contains information on how each schedule should be treated. It defines the behavior and constraint of the reservation (eg. refund period, reservation cost). Policies can be attached to a reservation schedule and businesses can own more than one policy.

```json
{
  "name": "String",
  "description": "String",
  "before": "32-bit Integer",
  "freeCancelDeadline": "32-bit Integer",
  "cancelRate": "Decimal",
  "basePrice": "Decimal"
}
```

- **name** - This uniquely identifies a policy within the array of other policy in the business object
- **description** - Short description of the policy
- **before** - How many days before the reservation will customers be able to start making reservation (eg. 7 states that customers can start making reservation 7 days before the actual date)
- **freeCancelDeadline** - Define how many days before the date can customers still cancel with full refund (eg. 2 states that any cancellation 2 days before the actual date will be fully refunded)
- **cancelRate** - This defines how much the customers will have to pay in percentage of the reservation cost
- **basePrice** - Base price for making a reservation

### `reservation`

The **reservation** collection contains information on a reservation schedule. A single reservation contains information on the date, time, placement (seats and availability) and menu (for pre-order).

```json
{
  "_id": "ObjectId [UNIQUE] [COMPOSITE 1]",
  "businessId": "ObjectId [COMPOSITE 1]",
  "name": "String",
  "start": "Date",
  "end": "Date",
  "placement": "placement",
  "menu_item": "Array<menu_item>",
  "policy": "policy"
}
```

- **_id** - This is MongoDB default uniquely generated id
- **businessId** - This is a parent reference to the business
- **name** - The name of the reservation (backend should handle setting a default if not specified)
- **date** - Date and time period of the reservation
- **placement** - Placement document object
- **menu_item** - List of items (not to be confused with menu document type)
- **policy** - Policy document object

### `order`

The **order** collection contains information on a each reservation order made.

```json
{
  "_id": "ObjectId",
  "customerId": "ObjectId",
  "businessId": "ObjectId",
  "paymentDate": "Date",
  "start": "Date",
  "end": "Date",
  "reserve": "Array<String>",
  "item": "Array<menu_item>",
  "basePrice": "Decimal",
  "totalPrice": "Decimal",
  "status": "String",
}
```

- **_id** - This is MongoDB default uniquely generated id
- **customerId** - This is a parent reference to the customer
- **businessId** - This a parent reference to the business
- **paymentDate** - When the customer paid for the reservation
- **start** - The start date and time of the reservation
- **end** - The end date and time of the reservation
- **reserve** - Array of reserved seat/tables id
- **item** - Array of items pre-ordered from the menu
- **basePrice** - The base price defined in the policy for the reservation
- **totalPrice** - The total cost of making the reservation including the items
- **status** - The status of the order, can be paid, used, expired or cancelled

## **Scripts**

This section contains the list of scripts to interact with out database during setup, testing, deployment and maintenance.
