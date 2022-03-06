from django.contrib import admin
from .models import User, Task


class UserAdmin(admin.ModelAdmin):
    list_display = (
        "username",
        "email",
        "is_staff",
    )


class TaskAdmin(admin.ModelAdmin):
    list_display = (
        "title",
        "points",
    )


admin.site.register(User, UserAdmin)
admin.site.register(Task, TaskAdmin)
