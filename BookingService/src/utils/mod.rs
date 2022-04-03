use postgres::{Client, Error, NoTls};
use chrono::{prelude::*, Duration};
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize)]
pub struct Reservation {
    id: i32,
    user_id: i32,
    room_id: i32,
    cancelled: bool,
    start: String,
    end: String,
    username: String
}

#[derive(Serialize, Deserialize)]
pub struct ReservationCreateRequest {
    user_id: i32,
    room_id: i32,
    start: String,
    end: String,
    username: String
}

#[derive(Serialize)]
struct Response {
    message: String
}

pub fn seed_db() -> Result<(), Error> {
    let mut client = Client::connect(
        "postgresql://postgres:root@localhost:5432/BookingServiceDB",
        NoTls,
    )?;

    client.batch_execute("DROP TABLE IF EXISTS reservations")?;

    client.batch_execute(
        "
        CREATE TABLE IF NOT EXISTS reservations (
            id              SERIAL PRIMARY KEY,
            userId          INTEGER,
            roomId          INTEGER,
            cancelled        BOOLEAN,
            reservation_start VARCHAR NOT NULL,
            reservation_end VARCHAR NOT NULL,
            user_username VARCHAR NOT NULL
            )
    ",
    )?;

    client.execute(
        "INSERT INTO reservations (userId, roomId, cancelled, reservation_start, reservation_end, user_username) VALUES ($1, $2, $3, $4, $5, $6)",
        &[&2.to_owned(), &1.to_owned(), &false, &str::replace(&Utc.ymd(2022, 3, 5).to_string(), "UTC", ""), &str::replace(&Utc.ymd(2022, 3, 7).to_string(), "UTC", ""), &"makulica"],
    )?;

    client.execute(
        "INSERT INTO reservations (userId, roomId, cancelled, reservation_start, reservation_end, user_username) VALUES ($1, $2, $3, $4, $5, $6)",
        &[&2.to_owned(), &2.to_owned(), &false, &str::replace(&Utc.ymd(2022, 4, 6).to_string(), "UTC", ""), &str::replace(&Utc.ymd(2022, 4, 7).to_string(), "UTC", ""), &"makulica"],
    )?;

    client.execute(
        "INSERT INTO reservations (userId, roomId, cancelled, reservation_start, reservation_end, user_username) VALUES ($1, $2, $3, $4, $5, $6)",
        &[&2.to_owned(), &1.to_owned(), &true, &str::replace(&Utc.ymd(2022, 3, 31).to_string(), "UTC", ""), &str::replace(&Utc.ymd(2022, 4, 7).to_string(), "UTC", ""), &"makulica"],
    )?;

    client.execute(
        "INSERT INTO reservations (userId, roomId, cancelled, reservation_start, reservation_end, user_username) VALUES ($1, $2, $3, $4, $5, $6)",
        &[&3.to_owned(), &1.to_owned(), &false, &str::replace(&Utc.ymd(2022, 4, 9).to_string(), "UTC", ""), &str::replace(&Utc.ymd(2022, 4, 15).to_string(), "UTC", ""), &"skafiskafnjak"],
    )?;

    client.close()?;

    Ok(())
}

pub fn get_all_for_room(id: i32) -> Result<(String), Error> {
    let mut client = Client::connect(
        "postgresql://postgres:root@localhost:5432/BookingServiceDB",
        NoTls,
    )?;

    let mut ret: Vec<Reservation> = vec![];
    for row in client.query("SELECT id, userId, roomId, cancelled, reservation_start, reservation_end, user_username FROM reservations WHERE roomId = $1 and cancelled = false", &[&id])? {
        let id: i32 = row.get(0);
        let user_id: i32 = row.get(1);
        let room_id: i32 = row.get(2);
        let cancelled: bool = row.get(3);
        let start: &str = row.get(4);
        let end: &str = row.get(5);
        let end_date = NaiveDate::parse_from_str(end, "%Y-%m-%d").unwrap();
        let username: &str = row.get(6);
        let today = Utc::now().naive_local().date();
        if end_date >= today {
            ret.push(Reservation { id: (id), user_id: (user_id), room_id: (room_id), cancelled: (cancelled), start: (start.to_string()), end: (end.to_string()), username: (username.to_string()) });
        }
    }
    client.close()?;

    Ok(serde_json::to_string(&ret).unwrap())
}

pub fn get_all_for_user(id: i32) -> Result<(String), Error> {
    let mut client = Client::connect(
        "postgresql://postgres:root@localhost:5432/BookingServiceDB",
        NoTls,
    )?;

    let mut ret: Vec<Reservation> = vec![];
    for row in client.query("SELECT id, userId, roomId, cancelled, reservation_start, reservation_end, user_username FROM reservations WHERE userId = $1", &[&id])? {
        let id: i32 = row.get(0);
        let user_id: i32 = row.get(1);
        let room_id: i32 = row.get(2);
        let cancelled: bool = row.get(3);
        let start: &str = row.get(4);
        let end: &str = row.get(5);
        let start_date = NaiveDate::parse_from_str(start, "%Y-%m-%d").unwrap();
        let end_date = NaiveDate::parse_from_str(end, "%Y-%m-%d").unwrap();
        let username: &str = row.get(6);
        ret.push(Reservation { id: (id), user_id: (user_id), room_id: (room_id), cancelled: (cancelled), start: (start.to_string()), end: (end.to_string()), username: (username.to_string()) });
    }
    client.close()?;

    Ok(serde_json::to_string(&ret).unwrap())
}

pub fn cancel_reservation(id: i32) -> Result<(String), Error> {
    let mut client = Client::connect(
        "postgresql://postgres:root@localhost:5432/BookingServiceDB",
        NoTls,
    )?;

    for row in client.query("SELECT reservation_start FROM reservations WHERE id = $1", &[&id])? {
        let start: &str = row.get(0);
        let start_date = NaiveDate::parse_from_str(start, "%Y-%m-%d").unwrap();
        let day_after_tomorrow = (Utc::now() + Duration::days(2)).naive_local().date();
        if day_after_tomorrow > start_date {
            client.close()?;
            return Ok(serde_json::to_string(&Response{message: "Cannot cancel a reservation that is in less than 48h.".to_string()}).unwrap());
        }
    }

    client.execute(
        "UPDATE reservations SET cancelled=true WHERE id=$1",
        &[&id],
    )?;

    client.close()?;

    Ok(serde_json::to_string(&Response{message: "OK".to_string()}).unwrap())
}

pub fn create_reservation(reservation: rocket::serde::json::Json<ReservationCreateRequest>) -> Result<(String), Error> {
    let mut client = Client::connect(
        "postgresql://postgres:root@localhost:5432/BookingServiceDB",
        NoTls,
    )?;

    let start_date_new = NaiveDate::parse_from_str(reservation.start.as_str(), "%Y-%m-%d").unwrap();
    let end_date_new = NaiveDate::parse_from_str(reservation.end.as_str(), "%Y-%m-%d").unwrap();
    for row in client.query("SELECT reservation_start, reservation_end FROM reservations WHERE roomId = $1 and cancelled = false", &[&reservation.room_id])? {
        let start: &str = row.get(0);
        let end: &str = row.get(1);
        let start_date_row = NaiveDate::parse_from_str(start, "%Y-%m-%d").unwrap();
        let end_date_row = NaiveDate::parse_from_str(end, "%Y-%m-%d").unwrap();
        if (start_date_new < end_date_row) && (end_date_new > start_date_row) {
            client.close()?;
            return Ok(serde_json::to_string(&Response{message: "Reservation overlaps with another reservation.".to_string()}).unwrap());
        }
        if end_date_new <= start_date_new {
            client.close()?;
            return Ok(serde_json::to_string(&Response{message: "Invalid reservation period.".to_string()}).unwrap());
        }
    }
    
    client.execute(
        "INSERT INTO reservations (userId, roomId, cancelled, reservation_start, reservation_end, user_username) VALUES ($1, $2, $3, $4, $5, $6)",
        &[&reservation.user_id, &reservation.room_id, &false, &start_date_new.to_string(), &end_date_new.to_string(), &reservation.username],
    )?;

    client.close()?;

    Ok(serde_json::to_string(&Response{message: "OK".to_string()}).unwrap())
}
