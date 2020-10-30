# Database

- [Database](#database)
  - [**Overview**](#overview)
  - [**MongoDB Schema**](#mongodb-schema)
    - [`customer`](#customer)
    - [`owner`](#owner)
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
  dob: String,
  avatar: String,
  preference: Array<String>,
  favorite: Array<String>,
}
```

- **_id** - This is MongoDB default uniquely generated id
- **name** - The name of the user, is used on authentication
- **password** - Hashed password in string format
- **dob** - Date of birth of the user
- **avatar** - Link to the avatar image asset of the user
- **preferences** - List of ids associated with a type of business
- **favorites** - List of ids associated with a business

### `owner`

The **owner** collection contains information on *business users*, these are users that owns a business and can configure their business store through our web application.

```json
{
  _id: ObjectId [UNIQUE],
  name: String [UNIQUE],
  password: String,
  avatar: String,
  businesses: Array<String>
}
```

- **_id** - This is MongoDB default uniquely generated id
- **name** - The name of the user, is used on authentication
- **password** - Hashed password in string format
- **avatar** - Link to the avatar image asset of the user
- **businesses** - List of ids associated with to the user's businesses

### `business`

The **business** collection contains information on the businesses owned by *business users*.

### `placement`

The **placement** collection contains information on the *floor layout* of a business. Businesses can have several placement that they can choose to apply to their reservation schedule.

### `menu`

The **menu** collection contains a *list of items** which can be customers can browse through or pre-order during reservation. Businesses can have several menu that they choose to apply to their reservation schedule.

### `menu_item`

The **menu_item** collection contains information on a particular item.

### `reservation`

The **reservation** collection contains information on a reservation schedule. A single reservation contains information on the date, time, placement (seats and availability) and menu (for pre-order).

### `policy`

The **policy** collection contains information on how each schedule should be treated. Policies can be attached to a reservation schedule and businesses can own more than one policy.

### `order`

The **order** collection contains information on a each reservation order made.

## **Scripts**

This section contains the list of scripts to interact with out database during setup, testing, deployment and maintenance.
