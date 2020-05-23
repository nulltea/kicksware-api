function mapInit() {
	let homeCoordinates = new google.maps.LatLng(50.416101, 30.6324291);
	let workCoordinates = new google.maps.LatLng(52.5080931, 13.4505881);
	let map = new google.maps.Map(document.getElementById('map'), {
		center: workCoordinates,
		zoom: 12,
		styles: $("body").hasClass("theme-dark") ? mapDarkTheme() : mapLightTheme(),
	});

	google.maps.event.addDomListener(window, 'resize', function () {
		setTimeout(function () {
			google.maps.event.trigger(map, "resize");
			map.setCenter(workCoordinates);
		}, 1400);
	});

	window.addEventListener("theme-change", function (event) {
		map.setOptions({styles: event.detail.theme === "dark" ? mapDarkTheme() : mapLightTheme()});
	})
}

function mapLightTheme() {
	return [
		{
			"featureType": "administrative",
			"elementType": "labels.text.fill",
			"stylers": [
				{
					"color": "#444444"
				}
			]
		},
		{
			"featureType": "landscape",
			"elementType": "all",
			"stylers": [
				{
					"color": "#f2f2f2"
				}
			]
		},
		{
			"featureType": "road",
			"elementType": "all",
			"stylers": [
				{
					"saturation": -100
				},
				{
					"lightness": 45
				}
			]
		},
		{
			"featureType": "road.highway",
			"elementType": "all",
			"stylers": [
				{
					"visibility": "simplified"
				}
			]
		},
		{
			"featureType": "water",
			"elementType": "all",
			"stylers": [
				{
					"color": "#46bcec"
				},
				{
					"visibility": "on"
				}
			]
		},
		{
			"featureType": "water",
			"elementType": "geometry",
			"stylers": [
				{
					"color": "#00b7bd"
				}
			]
		}
	];
}

function mapDarkTheme() {
	return [
		{
			"elementType": "geometry",
			"stylers": [
				{
					"color": "#1d2c4d",
				},
			],
		},
		{
			"elementType": "labels.text.fill",
			"stylers": [
				{
					"color": "#8ec3b9",
				},
			],
		},
		{
			"elementType": "labels.text.stroke",
			"stylers": [
				{
					"color": "#1a3646",
				},
			],
		},
		{
			"featureType": "administrative.country",
			"elementType": "geometry.stroke",
			"stylers": [
				{
					"color": "#4b6878",
				},
			],
		},
		{
			"featureType": "administrative.land_parcel",
			"elementType": "labels.text.fill",
			"stylers": [
				{
					"color": "#64779e",
				},
			],
		},
		{
			"featureType": "administrative.province",
			"elementType": "geometry.stroke",
			"stylers": [
				{
					"color": "#4b6878",
				},
			],
		},
		{
			"featureType": "landscape.man_made",
			"elementType": "geometry.stroke",
			"stylers": [
				{
					"color": "#334e87",
				},
			],
		},
		{
			"featureType": "landscape.natural",
			"elementType": "geometry",
			"stylers": [
				{
					"color": "#023e58",
				},
			],
		},
		{
			"featureType": "poi",
			"elementType": "geometry",
			"stylers": [
				{
					"color": "#283d6a",
				},
			],
		},
		{
			"featureType": "poi",
			"elementType": "labels.text.fill",
			"stylers": [
				{
					"color": "#6f9ba5",
				},
			],
		},
		{
			"featureType": "poi",
			"elementType": "labels.text.stroke",
			"stylers": [
				{
					"color": "#1d2c4d",
				},
			],
		},
		{
			"featureType": "poi.park",
			"elementType": "geometry.fill",
			"stylers": [
				{
					"color": "#023e58",
				},
			],
		},
		{
			"featureType": "poi.park",
			"elementType": "labels.text.fill",
			"stylers": [
				{
					"color": "#3C7680",
				},
			],
		},
		{
			"featureType": "road",
			"elementType": "geometry",
			"stylers": [
				{
					"color": "#304a7d",
				},
			],
		},
		{
			"featureType": "road",
			"elementType": "labels.text.fill",
			"stylers": [
				{
					"color": "#98a5be",
				},
			],
		},
		{
			"featureType": "road",
			"elementType": "labels.text.stroke",
			"stylers": [
				{
					"color": "#1d2c4d",
				},
			],
		},
		{
			"featureType": "road.highway",
			"elementType": "geometry",
			"stylers": [
				{
					"color": "#2c6675",
				},
			],
		},
		{
			"featureType": "road.highway",
			"elementType": "geometry.stroke",
			"stylers": [
				{
					"color": "#255763",
				},
			],
		},
		{
			"featureType": "road.highway",
			"elementType": "labels.text.fill",
			"stylers": [
				{
					"color": "#b0d5ce",
				},
			],
		},
		{
			"featureType": "road.highway",
			"elementType": "labels.text.stroke",
			"stylers": [
				{
					"color": "#023e58",
				},
			],
		},
		{
			"featureType": "transit",
			"elementType": "labels.text.fill",
			"stylers": [
				{
					"color": "#98a5be",
				},
			],
		},
		{
			"featureType": "transit",
			"elementType": "labels.text.stroke",
			"stylers": [
				{
					"color": "#1d2c4d",
				},
			],
		},
		{
			"featureType": "transit.line",
			"elementType": "geometry.fill",
			"stylers": [
				{
					"color": "#283d6a",
				},
			],
		},
		{
			"featureType": "transit.station",
			"elementType": "geometry",
			"stylers": [
				{
					"color": "#3a4762",
				},
			],
		},
		{
			"featureType": "water",
			"elementType": "geometry",
			"stylers": [
				{
					"color": "#0e1626",
				},
			],
		},
		{
			"featureType": "water",
			"elementType": "labels.text.fill",
			"stylers": [
				{
					"color": "#4e6d70",
				},
			],
		},
	];
}

$(document).ready(function () {
	"use strict";
	mapInit();
});