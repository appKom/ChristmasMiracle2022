from django.db import models
from django.contrib.auth.models import AbstractBaseUser
from django.contrib.auth.models import PermissionsMixin
from django.contrib.auth.models import BaseUserManager


class UserProfileManager(BaseUserManager):
    # Create a new user profile
    def create_user(self, nickname, password, owID):

        if not owID:
            raise ValueError("Must have an owID")

        user = self.model(nickname=nickname, owID=owID)

        user.set_password(password)
        user.save(using=self._db)

        return user

    # Create a new superuser profile
    def create_superuser(self, nickname, password, owID):

        user = self.create_user(nickname, password, owID)
        user.is_superuser = True
        user.is_staff = True

        user.save(using=self._db)

        return user


# Database model for users in the system
class User(AbstractBaseUser, PermissionsMixin):

    # Fields for the user model
    owID = models.IntegerField(primary_key=True)
    family_name = models.CharField(max_length=255, blank=True, null=True)
    given_name = models.CharField(max_length=255, blank=True, null=True)
    name = models.CharField(max_length=255, blank=True, null=True)
    nickname = models.CharField(max_length=255, unique=True)
    picture = models.CharField(max_length=1000, blank=True, null=True)
    preferred_username = models.CharField(
        max_length=255, blank=True, null=True)
    is_staff = models.BooleanField(default=False)

    points = models.IntegerField(default=0)
    flags = models.IntegerField(default=0)

    # Usermanager for creating users
    objects = UserProfileManager()

    USERNAME_FIELD = "nickname"
    REQUIRED_FIELDS = ["owID"]

    def __str__(self):
        return self.nickname


class Task(models.Model):
    # Fields for the task model
    title = models.CharField(max_length=255)
    description = models.CharField(max_length=255)
    points = models.IntegerField(default=0)

    def __str__(self):
        return self.task_name
