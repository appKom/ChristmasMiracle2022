from .models import User
from rest_framework import serializers


class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = User
        fields = ('password', 'owID', 'family_name', 'given_name', 'name', 'nickname', 'picture', 'preferred_username', 'is_staff', 'points', 'flags')
        extra_kwargs = {'password': {'write_only': True}}

    def create(self, validated_data):
        user = User(
            nickname=validated_data['nickname'],
            owID = validated_data['owID'],
        )

        user.set_password(validated_data['password'])
        user.save()
        return user
    
    def update(self, instance, validated_data):
        instance.given_name = validated_data.get('given_name', instance.given_name)

        instance.save()
        return instance