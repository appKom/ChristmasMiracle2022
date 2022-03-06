from django.shortcuts import render

from core.models import User
from rest_framework.generics import CreateAPIView, RetrieveAPIView
from django.views import View
from django.http import HttpResponse
from rest_framework.permissions import AllowAny, IsAuthenticated
from rest_framework.authtoken.models import Token
from core.serializers import UserSerializer
from rest_framework.authtoken.views import ObtainAuthToken
from django.contrib.auth import authenticate
from rest_framework.response import Response


class LoginView(ObtainAuthToken):
    """
    Loginview for obtaining token
    """

    def post(self, request, *args, **kwargs):
        try:
            user = authenticate(
                request=request,
                username=request.data["username"],
                password=request.data["password"],
            )
            if not user:
                return Response(status=400)
            token, created = Token.objects.get_or_create(user=user)
            return Response(
                {
                    "token": token.key,
                    "user": UserSerializer(user).data,

                }
            )
        except:
            return Response(status=400)


class ValidateTokenView(RetrieveAPIView):
    permission_classes = (IsAuthenticated,)

    def get(self, request, *args, **kwargs):
        return HttpResponse(status=200)


class RegisterView(CreateAPIView):
    queryset = User.objects.all()
    permission_classes = (AllowAny,)
    serializer_class = UserSerializer

    def post(self, request, *args, **kwargs):
        response = super().post(request, *args, **kwargs)

        token, created = Token.objects.get_or_create(
            user_id=response.data["pk"])
        response.data["token"] = str(token)
        return response
