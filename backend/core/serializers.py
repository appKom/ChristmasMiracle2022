from .models import User
from rest_framework import serializers


class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = ('password', 'username', 'email', 'is_staff', 'solvedTasks')
        extra_kwargs = {'password': {'write_only': True}}

    def create(self, validated_data):
        user = User(
            nickname=validated_data['username'],
        )

        user.set_password(validated_data['password'])
        user.save()
        return user

    def update(self, instance, validated_data):
        instance.save()
        return instance
