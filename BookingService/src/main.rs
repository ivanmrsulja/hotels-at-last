#[macro_use] 
extern crate rocket;

mod utils;

use rocket::{routes, Rocket, Build};
use rocket::response::content::{self};
use utils::ReservationCreateRequest;
use std::{thread};

#[get("/<id>")]
fn get_all_for_room(id: i32) -> content::Json<String> {
    thread::spawn(move || {
        content::Json(utils::get_all_for_room(id).unwrap())
    }).join().unwrap()
}

#[get("/user/<id>")]
fn get_all_for_user(id: i32) -> content::Json<String> {
    thread::spawn(move || {
        content::Json(utils::get_all_for_user(id).unwrap())
    }).join().unwrap()
}

#[put("/<id>/cancel")]
fn cancel_reservation(id: i32) -> content::Json<String> {
    thread::spawn(move || {
        content::Json(utils::cancel_reservation(id).unwrap())
    }).join().unwrap()
}

#[post("/", format = "json", data = "<reservation>")]
fn create_reservation(reservation: rocket::serde::json::Json<ReservationCreateRequest>) -> content::Json<String> {
    thread::spawn(move || {
        content::Json(utils::create_reservation(reservation).unwrap())
    }).join().unwrap()
}

#[launch]
fn rocket() -> Rocket<Build> {
    thread::spawn(|| {
        utils::seed_db();
    }).join().expect("Thread panicked");

    rocket::build().mount("/api/reservations", routes![get_all_for_room, get_all_for_user, cancel_reservation, create_reservation])
}
