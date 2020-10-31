# Database

- [Database](#database)
  - [**Overview**](#overview)
  - [**MongoDB Schema**](#mongodb-schema)
    - [`customer`](#customer)
    - [`business`](#business)
      - [`placement`](#placement)
      - [`menu`](#menu)
      - [`menu_item`](#menu_item)
    - [`reservation`](#reservation)
    - [`policy`](#policy)
    - [`order`](#order)
  - [**Scripts**](#scripts)

## **Overview**

This repository contains files related to creating, testing, deploying and maintaining Seatlect's databases. Seatlect uses MongoDB as its primary data store for structured data.

## **MongoDB Schema**

This section contains the schema for MongoDB database. Each one corresponds to a single collection.

### `customer`

The **customer** collection contains information on *general users*, these are users which does own a business. They are the customer of a business. The following defines the schema for the document in the collection.

```json
{
  _id: ObjectId [UNIQUE],
  name: String [UNIQUE],
  password: String,
  dob: Date,
  avatar: String,
  preference: Array<String>,
  favorite: Array<ObjectId>
}
```

- **_id** - This is MongoDB default uniquely generated id
- **name** - The name of the user, is used on authentication
- **password** - Hashed password in string format
- **dob** - Date of birth of the user
- **avatar** - Link to the avatar image asset of the user
- **preferences** - List of ids associated with a type of business
- **favorites** - List of ids associated with a business

### `business`

The **business** collection contains information on *business users* and there business, these are users that owns a business and can configure their business store through our web application. The following defines the schema for the document in the collection.

```json
{
  _id: ObjectId [UNIQUE],
  name: String [UNIQUE],
  password: String,
  businessName: String,
  type: String,
  description: String,
  displayImage: String,
  images: Array<String>,
  placement: Array<placement>,
  menu: Array<menu>,
  menu_item: Array<menu_item>
}
```

- **_id** - This is MongoDB default uniquely generated id
- **name** - The name of the user, is used on authentication
- **password** - Hashed password in string format
- **businessName** - The name of the business, is not unique
- **type** - The type associated with this business
- **description** - Short description of the business, will be displayed on the mobile application
- **displayImage** - Equivalent to a profile photo
- **images** - List of image resource url
- **placement** - Array of placement document objects
- **menu** -Array of menu document objects
- **menu_item** - Array of menu_item document

Note that placement and menu are **templates** which is used during the **creation of reservation document**. Templates are used as a base for information to be contained in a reservation. They can then be **modified** before finalization of a reservation object creation.

#### `placement`

The **placement** document contains information on the *floor layout template* of a business. Businesses can have several placement that they can choose to apply to their reservation schedule. The following define the schema for a placement document.

**`TODO`** *Define schema and detailed description of schema fields*

```json
{
  name: String,
  default: Boolean
}
```

- **name** - This uniquely identifies a placement template within the array of other placement template in the business object
- **default** - If true indicates that this placement should be on display in the mobile application

#### `menu`

The **menu** document contains a *list of menu_item names* of a business. Businesses can have several menu that they choose to apply to their reservation schedule. The following define the schema for a placement document.

```json
{
  name: String,
  description: String,
  items: Array<String>,
  default: Boolean
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
  name: String,
  description: String,
  image: String,
  price: Double
}
```

- **name** - This uniquely identifies an item within the array of other menu_items in the business object
- **description** - Short description of the item
- **image** - Image resource url
- **price** - The price of a single order of this item

### `reservation`

The **reservation** collection contains information on a reservation schedule. A single reservation contains information on the date, time, placement (seats and availability) and menu (for pre-order).

```json
{
  _id: ObjectId,
  businessId: ObjectId,
  name: String,
  date: Date,
  placement: placement,
  menu_item: Array<menu_item>
}
```

- **_id** - This is MongoDB default uniquely generated id
- **name** - The name of the reservation (backend should handle setting a default if not specified)
- **date** - Date and time period of the reservation
- **placement** - Placement document object
- **menu_item** - List of items (not to be confused with menu document type)

### `policy`

The **policy** collection contains information on how each schedule should be treated. Policies can be attached to a reservation schedule and businesses can own more than one policy.

### `order`

The **order** collection contains information on a each reservation order made.

## **Scripts**

This section contains the list of scripts to interact with out database during setup, testing, deployment and maintenance.
