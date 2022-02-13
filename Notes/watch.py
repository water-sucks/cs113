"""
Watch notes and rebuild PDFs if they are modified or created
"""

import sys
import subprocess
import time
from watchdog.observers import Observer
from watchdog.events import FileClosedEvent, PatternMatchingEventHandler


class Watch:
    def __init__(self, directory: str) -> None:
        self.directory = directory 
        self.observer = Observer()

    def run(self):
        event_handler = Handler(patterns=["*.md"])
        self.observer.schedule(event_handler, self.directory, recursive=True)
        self.observer.start()
        try:
            while True:
                time.sleep(5)
        except:
            print("Stopping observer")
        finally:
            self.observer.stop()
        self.observer.join()


class Handler(PatternMatchingEventHandler):
    def __init__(self, patterns=None, ignore_patterns=None, ignore_directories=True, case_sensitive=True):
        super().__init__(patterns, ignore_patterns, ignore_directories, case_sensitive)

    @staticmethod
    def on_closed(event: FileClosedEvent) -> None:
        if event.is_directory or event.is_synthetic:
            return
        destination = event.src_path.replace(".md", ".pdf")
        subprocess.run(["pandoc", "-f", "markdown", event.src_path, "-o", destination]) 
        

if __name__ == "__main__":
    watch = Watch(sys.argv[1] if len(sys.argv) > 1 else "./")
    watch.run()
