"""
Find Source Code Disclosure Vulnerability by exposed .git/
"""
import requests
from threading import Thread
import argparse
import time

class Color:        
    class Fore:
        RED = '\x1b[31m'   
        WHITE = '\x1b[37m'
        YELLOW = '\x1b[33m'
        GREEN = '\x1b[32m'
    class Back:
        YELLOW = '\x1b[43m'
    class Style:
        RESET_ALL = '\x1b[0m'
            


BANNER = """
        ___      ___________     __________
        \  \    /  /  ____  \ __|   _______|
         \  \  /  /  /    \__|__|  |
          \  \/  /  /   _____ __|  |_____
           |    |  |   |__   |  |   _____|
          /  /\  \  \     |  |  |  |
         /  /  \  \  \____|  |  |  |
        /__/    \__\_________|__|__|
        
        # Exposed Git Finder by Prasant 
        (https://github.com/prasant_paudel/xgif.git)
"""

def parse_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('-d', '--domain', help="Domain name to check for .git exposure")
    parser.add_argument('-D', '--domains', help="Domain names form a file")
    parser.add_argument('-t', '--threads', help="Number of threads to use (default=40)", type=int)
    # parser.add_argument('-o', '--output', help="Output file to save the results")
    parser.add_argument('-v', '--verbose', help="Enable Verbosity and display refused connections", nargs='?', default=False)
    return parser.parse_args()

def __check_url(url:str) -> str:
    try:
        with requests.get(url, allow_redirects=False) as resp:
            if resp.status_code == 200:
                if 'directory listing' in resp.text.lower() or 'index of' in resp.text.lower():
                    return f"{url} --> *** Potentially Exploitable ***"
                else:
                    return f"{url} --> Status: 200 (Check manually)"
            if resp.status_code == 403:
                return f"{url} --> 403 Forbidden (Exists but restricted)"
    except requests.exceptions.ConnectionError:
        return f"{url} --> Connection Error"
    return ''

def __display_resp(resp:str, verbose=False):
    if not resp:
        return None
    if 'Exploitable' in resp:
        print(f"{Color.Fore.RED}{Color.Back.YELLOW}{resp}{Color.Style.RESET_ALL}")
        time.sleep(0.1)
    elif 'Status: 200' in resp:
        print(f"{Color.Fore.GREEN}{resp}{Color.Style.RESET_ALL}")
        time.sleep(0.1)
    elif '403 Forbidden' in resp:
        print(f"{Color.Fore.YELLOW}{resp}{Color.Style.RESET_ALL}")
        time.sleep(0.1)
    elif verbose and 'Connection Error' in resp:
        print(f"{Color.Fore.RED}{resp}{Color.Style.RESET_ALL}")
        time.sleep(0.1)

def check_git_exposure(domain:str, verbose) -> str:
    # Check /.git
    __display_resp(__check_url(f"http://{domain.strip()}/.git"), verbose)
    # Check /.git/config
    __display_resp(__check_url(f"http://{domain.strip()}/.git/config"), verbose)
    # Check /.git/HEAD
    __display_resp(__check_url(f"http://{domain.strip()}/.git/HEAD"), verbose)
    # Check /.git/logs/HEAD
    __display_resp(__check_url(f"http://{domain.strip()}/.git/logs/HEAD"), verbose)
    # Check /.git/index
    __display_resp(__check_url(f"http://{domain.strip()}/.git/index"), verbose)

def chunkify(iterable, thread_count=40):
    chunksize = int(len(iterable) / thread_count)
    if chunksize <= 1:
        return [[_] for _ in iterable]
    return [iterable[_:_+chunksize] for _ in range(0, len(iterable), chunksize)]


def enum_domains(domains, verbose:bool=False):
    for domain in domains:
        time.sleep(0.1)
        resp = check_git_exposure(domain, verbose)
        __display_resp(resp)

def main():
    print(Color.Fore.RED + BANNER + Color.Style.RESET_ALL)
    args = parse_args()
    verbose = (args.verbose or args.verbose == None)
    domains = list()
    if args.domain:
        check_git_exposure(args.domain.split('://')[-1].strip().strip('/'), verbose)
        exit()
    elif args.domains:
        with open(args.domains) as f:
            domains = list(set([_.split('://')[-1].strip().strip('/') for _ in f.readlines() if _.strip()]))
    if not args.threads:
        args.threads = 40
    

    _threads = []
    chunks = tuple(chunkify(domains, args.threads))
    for chunk in chunks:
        thread = Thread(target=enum_domains, args=(chunk, verbose,), daemon=True)
        thread.start()
        _threads.append(thread)
    for thread in _threads[:-1]:
        thread.join()


if __name__ == '__main__':
    main()
    
        


