from django.shortcuts import render
from .models import Task
from .serializers import TaskSerializer
from rest_framework import viewsets
from rest_framework import permissions
# Create your views here.


class TaskViewSet(viewsets.ModelViewSet):
    queryset = Task.objects.all()
    serializer_class = TaskSerializer
    permission_classes = [permissions.IsAuthenticated]
