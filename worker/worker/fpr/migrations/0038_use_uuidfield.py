# Generated by Django 1.11.29 on 2023-08-11 16:29
import uuid

from django.db import migrations

from worker.main.models import UUIDField


class Migration(migrations.Migration):
    dependencies = [
        ("fpr", "0037_update_idtools"),
    ]

    operations = [
        migrations.AlterField(
            model_name="format",
            name="uuid",
            field=UUIDField(
                default=uuid.uuid4,
                editable=False,
                help_text="Unique identifier",
                unique=True,
            ),
        ),
        migrations.AlterField(
            model_name="formatgroup",
            name="uuid",
            field=UUIDField(
                default=uuid.uuid4,
                editable=False,
                help_text="Unique identifier",
                unique=True,
            ),
        ),
        migrations.AlterField(
            model_name="formatversion",
            name="uuid",
            field=UUIDField(
                default=uuid.uuid4,
                editable=False,
                help_text="Unique identifier",
                unique=True,
            ),
        ),
        migrations.AlterField(
            model_name="fpcommand",
            name="uuid",
            field=UUIDField(
                default=uuid.uuid4,
                editable=False,
                help_text="Unique identifier",
                unique=True,
            ),
        ),
        migrations.AlterField(
            model_name="fprule",
            name="uuid",
            field=UUIDField(
                default=uuid.uuid4,
                editable=False,
                help_text="Unique identifier",
                unique=True,
            ),
        ),
        migrations.AlterField(
            model_name="fptool",
            name="uuid",
            field=UUIDField(
                default=uuid.uuid4,
                editable=False,
                help_text="Unique identifier",
                unique=True,
            ),
        ),
        migrations.AlterField(
            model_name="idcommand",
            name="uuid",
            field=UUIDField(
                default=uuid.uuid4,
                editable=False,
                help_text="Unique identifier",
                unique=True,
            ),
        ),
        migrations.AlterField(
            model_name="idrule",
            name="uuid",
            field=UUIDField(
                default=uuid.uuid4,
                editable=False,
                help_text="Unique identifier",
                unique=True,
            ),
        ),
        migrations.AlterField(
            model_name="idtool",
            name="uuid",
            field=UUIDField(
                default=uuid.uuid4,
                editable=False,
                help_text="Unique identifier",
                unique=True,
            ),
        ),
    ]
