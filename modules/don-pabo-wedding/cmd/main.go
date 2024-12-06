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
    <title>Kushan & Pabo | Save the Date</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <script src="https://cdn.tailwindcss.com"></script>
		<link rel="icon" href="https://thumbs.dreamstime.com/b/wedding-rings-line-icon-engagement-rings-vector-illustration-isolated-white-jewelery-outline-style-design-designed-wedding-123936888.jpg" type="image/x-icon">
		<link rel="preconnect" href="https://fonts.googleapis.com">
		<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
		<link href="https://fonts.googleapis.com/css2?family=Imperial+Script&family=Lora:ital,wght@0,400..700;1,400..700&family=Roboto+Mono:ital,wght@0,100..700;1,100..700&display=swap" rel="stylesheet">
		<style>
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
<body class="bg-stone-300 text-xl min-h-screen flex items-center justify-center bg-cover">
    <div class="text-center backdrop-blur-sm bg-white/30 p-10">
			<img src="https://cdn-icons-png.flaticon.com/512/620/620812.png" alt="Don Pabo & Kushan" class="mx-auto w-12 mb-5">
			<p class="tracking-[.9em] text-center roboto-mono uppercase">Pabo &</p>
			<p class="tracking-[.9em] text-center roboto-mono uppercase mb-4">Kushan</p>
			<hr class="border-1 border-black w-1/4 mx-auto">

			<div class="my-5 roboto-mono">
				<p class="tracking-[.5em] roboto-mono mb-5 uppercase text-center">03.10.25</p>
				<p class="tracking-[.4em]">MELBOURNE</p>
				<p class="tracking-[.4em]">AUSTRALIA</p>
			</div>
			
			<hr class="border-1 border-black w-1/4 mx-auto">
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
