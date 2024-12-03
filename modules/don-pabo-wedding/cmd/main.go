package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const home = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Save the Date</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <script src="https://cdn.tailwindcss.com"></script>
		<link rel="preconnect" href="https://fonts.googleapis.com">
		<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
		<link href="https://fonts.googleapis.com/css2?family=Imperial+Script&family=Lora:ital,wght@0,400..700;1,400..700&family=Roboto+Mono:ital,wght@0,100..700;1,100..700&display=swap" rel="stylesheet">
		<style>
			.lora {
				font-family: "Lora", serif;
				font-optical-sizing: auto;
				font-weight: <weight>;
				font-style: normal;
			}
			.imperial-script-regular {
				font-family: "Imperial Script", cursive;
				font-weight: 400;
				font-style: normal;
			}
			.roboto-mono {
				font-family: "Roboto Mono", monospace;
				font-optical-sizing: auto;
				font-style: normal;
			}
		</style>
</head>
<body class="bg-white min-h-screen flex items-center justify-center mx-10">
    <div class="text-center">
			<img src="https://t3.ftcdn.net/jpg/06/30/39/20/360_F_630392023_r1HSnZhXzafJoWCoikSMq0GSh2hMN3wN.jpg" alt="Don Pabo & Kushan" class="mx-auto w-1/2 md:w-1/3 lg:w-1/4">
			<h1 class="tracking-[1em] lora mb-5">PABO & KUSHAN</h1>
			<hr class="border-1 border-black w-1/4 mx-auto">
			<div class="tracking-[.125em] my-5 roboto-mono">
				<h2 class="text-2xl font-light">03 • 10 • 25</h2>
				<h2 class="mt-2">MELBOURNE, AUSTRALIA</h2>
			</div>
			<hr class="border-1 border-black w-1/4 mx-auto">
			<h2 class="tracking-wide lora  my-5">
				Join us in celebrating happily ever after!<br/>Formal invitation to follow.
			</h2>
    </div>
</body>
</html>
`

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", func(c echo.Context) error {
		return c.HTML(200, home)
	})
	e.HideBanner = true

	e.Logger.Fatal(e.Start(":80"))
}
