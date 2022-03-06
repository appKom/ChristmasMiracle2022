from django.db import models
from django.contrib.auth.models import AbstractBaseUser
from django.contrib.auth.models import PermissionsMixin
from django.contrib.auth.models import BaseUserManager


class UserProfileManager(BaseUserManager):
    # Create a new user profile
    def create_user(self, username, email, password):

        if not username:
            raise ValueError("Must have an username")

        if not email:
            raise ValueError("Must have an email")

        user = self.model(username=username, email=email)

        user.set_password(password)
        user.save(using=self._db)

        return user

    # Create a new superuser profile
    def create_superuser(self, username, email, password):

        user = self.create_user(username, email,  password)
        user.is_superuser = True
        user.is_staff = True

        user.save(using=self._db)

        return user


# Database model for users in the system
class User(AbstractBaseUser, PermissionsMixin):

    # Fields for the user model
    username = models.CharField(max_length=255, unique=True)
    email = models.CharField(max_length=255, blank=True, null=True)

    is_staff = models.BooleanField(default=False)

    solvedTasks = models.ManyToManyField(
        'Task', related_name='solvedTasks', blank=True)

    # Usermanager for creating users
    objects = UserProfileManager()

    USERNAME_FIELD = "username"
    REQUIRED_FIELDS = ["email"]

    def __str__(self):
        return self.username


class Task(models.Model):
    # Fields for the task model
    title = models.CharField(max_length=255)
    description = models.CharField(max_length=255)
    points = models.IntegerField(default=0)

    def __str__(self):
        return self.title
