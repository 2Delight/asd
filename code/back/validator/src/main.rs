use actix_web::{get, App, HttpResponse, HttpServer, Responder, Error};

#[get("/ping")]
async fn index() -> Result<impl Responder, Error> {
    Ok(HttpResponse::Ok().body("pong"))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(index)
    })
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}
