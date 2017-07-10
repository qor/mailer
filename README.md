# Mailer

Mail solution

## Usage

mailer := New(config) // Mailer

mailer.Send(Email, objs...)

mailer.Send(Email.Merge(Email), objs...)
