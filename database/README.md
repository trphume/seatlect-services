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
      - [`preorder`](#preorder)
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
  "username": "String [UNIQUE]",
  "password": "String",
  "dob": "Date",
  "avatar": "String",
  "preference": "Array<String>",
  "favorite": "Array<ObjectId>"
}
```

- **_id** - This is MongoDB default uniquely generated id
- **username** - The name of the user, is used on authentication
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
  "username": "String [UNIQUE]",
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
  "policy": "policy"
}
```

- **_id** - This is MongoDB default uniquely generated id
- **username** - The name of the user, is used on authentication
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
- **policy** - Business policy object

Note that placement and menu are **templates** which is used during the **creation of reservation document**. Templates are used as a base for information to be contained in a reservation. They can then be **modified** before finalization of a reservation object creation.

#### `placement`

The **placement** document contains information on the *floor layout template* of a business. Businesses can have several placement that they can choose to apply to their reservation schedule. The following define the schema for a placement document.

```json
{
  "name": "String",
  "entity": "Array<entity>",
}
```

Below is the **entity** document used in placement.

```json
{
  "name": "String",
  "floor": "32-bit Integer",
  "type": "String",
  "price": "Decimal",
  "user": "ObjectId",
  "status": "String",
  "x": "Double",
  "y": "Double"
}
```

- **name** - This uniquely identifies a placement template within the array of other placement template in the business object
- **entity** - An array of entity describing the placement layout
  - **name** - This has to be unique and is set upon creation of entity
  - **floor** - Indicates which floor the entity should be placed on
  - **type** - This indicates how the system should interpret the entity for example, TABLE or SEAT
  - **price** - The price of of reservation
  - **user** - Contain id of user who has reserve the seat
  - **status** - Contain the status of the entity - EMPTY,TAKEN,PROCESSING
  - **x** - x coords
  - **y** - y coords

#### `menu`

The **menu** document contains a *list of menu_item* of a business. Businesses can have several menu that they choose to apply to their reservation schedule. The following define the schema for a placement document.

```json
{
  "name": "String",
  "description": "String",
  "items": "Array<menu_item>",
  "default": "Boolean"
}
```

- **name** - This uniquely identifies a menu template within the array of other menu in the business object
- **description** - Short description of the menu template
- **items** - Array of menu_item objects
- **default** - If true indicates that this menu should be on display in the mobile application

#### `menu_item`

The **menu_item** document contains information on a particular item.

```json
{
  "name": "String",
  "description": "String",
  "image": "String",
  "price": "Decimal",
  "max": "32-bit Integer"
}
```

- **name** - This uniquely identifies an item within the array of other menu_items in the business object
- **description** - Short description of the item
- **image** - Image resource url
- **price** - The price of a single order of this item
- **max** - Max number of order per person

#### `policy`

The **policy** document contains information on a particular business.

```json
{
  "minAge": "32-bit Integer"
}
```

- **minAge** - Minimum age of a user that are allowed reserve a seat or table with the business

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
  "menu": "Array<menu_item>",
}
```

- **_id** - This is MongoDB default uniquely generated id
- **businessId** - This is a parent reference to the business
- **name** - The name of the reservation (backend should handle setting a default if not specified)
- **start** - Start time of the reservation
- **end** - End time of the reservation
- **placement** - Placement document object
- **menu** - List of items (not to be confused with menu document type)

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
  "reserve": "Array<entity>",
  "preorder": "Array<preorder>",
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
- **reserve** - Array of reserved seat/tables entity documents
- **preorder** - Array of items pre-ordered from the menu
- **totalPrice** - The total cost of making the reservation including the items
- **status** - The status of the order, can be paid, used, expired or cancelled

#### `preorder`

The **preorder** document contains information on a particular item and quantity ordered.

```json
{
  "name": "String",
  "quantity": "32-bit Integer",
  "price": "Decimal"
}
```

## **Scripts**

This section contains the list of scripts to interact with out database during setup, testing, deployment and maintenance.
