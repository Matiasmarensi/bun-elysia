import { Elysia, t } from "elysia";

const app = new Elysia()

  // Ruta GET en la raíz
  .get("/", () => "Welcome to the Elysia Server!")

  // Ruta con un parámetro dinámico
  .get("/hello/:name", ({ params }) => `Hello, ${params.name}!`)

  // Ruta POST con validación de datos
  .post(
    "/submit",
    ({ body }) => {
      return { message: "Data received successfully", received: body };
    },
    {
      // Validación de esquema para `body`
      body: t.Object({
        name: t.String(),
        age: t.Number(),
      }),
    }
  )

  // Ruta con Query Parameters
  .get("/search", ({ query }) => {
    const searchTerm = query.q || "nothing";
    return { message: `You searched for: ${searchTerm}` };
  })

  // Escuchar en el puerto 3000
  .listen(3000);

console.log("Server running at http://localhost:3000");
