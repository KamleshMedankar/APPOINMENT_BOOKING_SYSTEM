# Appointment Booking System (Golang + Gin + MySQL)

## 📌 Overview

This is a simple **Appointment Booking System** built using **Golang (Gin framework)** and **MySQL**.
It allows coaches to set availability and users to book 30-minute appointment slots.

---

## 🚀 Features

* Set coach availability (weekly schedule)
* View available time slots (30-minute intervals)
* Book appointments
* View user bookings
* Prevent double booking using database constraints

---

## 🛠️ Tech Stack

* Golang (Gin Framework)
* MySQL
* REST API

---

## 📂 API Endpoints

### 1️⃣ Set Coach Availability

**POST** `/coaches/:coach_id/availability`

#### Request Body:

```json
{
  "day_of_week": "monday",
  "start_time": "09:00",
  "end_time": "16:00"
}
```

---

### 2️⃣ Get Available Slots

**GET** `/coaches/:coach_id/slots?date=YYYY-MM-DD`

#### Example:

```
GET /coaches/2/slots?date=2026-04-06
```

#### Response:

```json
[
  "2026-04-06T09:00:00Z",
  "2026-04-06T09:30:00Z"
]
```

---

### 3️⃣ Book Appointment

**POST** `/users/bookings`

#### Request Body:

```json
{
  "user_id": 2,
  "coach_id": 2,
  "datetime": "2026-04-06T09:00:00Z"
}
```

---

### 4️⃣ Get User Bookings

**GET** `/users/:user_id/bookings`

#### Example:

```
GET /users/1/bookings
```

---

## 📮 Postman Requests

### Set Availability

```
POST http://localhost:3030/coaches/1/availability
```

Body:

```json
{
  "day_of_week": "monday",
  "start_time": "09:00",
  "end_time": "16:00"
}
```

---

### Get Slots

```
GET http://localhost:3030/coaches/2/slots?date=2026-04-06
```

---

### Book Appointment

```
POST http://localhost:3030/users/bookings
```

Body:

```json
{
  "user_id": 2,
  "coach_id": 2,
  "datetime": "2026-04-06T09:00:00Z"
}
```

---

### Get User Bookings

```
GET http://localhost:3030/users/1/bookings
```

---

## 🗄️ Database Schema

### Coaches

* id
* name

### Availability

* coach_id
* day_of_week
* start_time
* end_time

### Bookings

* user_id
* coach_id
* booking_datetime

---

## ⚡ How to Run

```bash
git clone <repo-url>
cd APPOINTMENT_BOOKING_SYSTEM
go mod tidy
go run main.go
```

Server runs on:

```
http://localhost:3030
```

---

## 🔒 Concurrency Handling

* Prevents double booking using:

```
UNIQUE (coach_id, booking_datetime)
```

---

## ⚠️ Assumptions

* All times are in IST (no timezone handling)
* Slots are generated in 30-minute intervals

---

## 📌 Future Improvements

* Timezone support (UTC conversion)
* Booking cancellation
* Authentication (JWT)
* Pagination for bookings

---

## 👨‍💻 Author

Kamlesh Medankar
