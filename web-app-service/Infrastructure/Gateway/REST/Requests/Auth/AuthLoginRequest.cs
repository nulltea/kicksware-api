
using System;
using System.Net;
using System.Security;
using Core.Entities.Users;
using RestSharp;

namespace Infrastructure.Gateway.REST.Auth
{
	public class AuthLoginRequest : AuthBaseRequest
	{
		public AuthLoginRequest(AuthCredentials credentials) : base("/login", Method.POST)
		{
			AddJsonBody(credentials);
		}
	}
}